/*
 * Copyright 2018 The Service Manager Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package postgres

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/Peripli/service-manager/pkg/util"

	"github.com/Peripli/service-manager/storage"
	"github.com/lib/pq"

	"github.com/Peripli/service-manager/pkg/types"

	"github.com/Peripli/service-manager/pkg/log"
	notificationConnection "github.com/Peripli/service-manager/storage/postgres/notification_connection"
)

const (
	postgresChannel       = "notifications"
	aTrue           int32 = 1
	aFalse          int32 = 0
)

type Notificator struct {
	isConnected int32
	isListening int32

	queueSize int

	connectionMutex *sync.Mutex
	connection      notificationConnection.NotificationConnection

	consumersMutex    *sync.Mutex
	consumers         *consumers
	storage           notificationStorage
	connectionCreator notificationConnectionCreator

	notificationFilters []storage.NotificationFilterFunc

	ctx context.Context

	lastKnownRevision int64
}

// NewNotificator returns new Notificator based on a given NotificatorStorage and desired queue size
func NewNotificator(st storage.Storage, settings *storage.Settings) (*Notificator, error) {
	ns, err := NewNotificationStorage(st)
	connectionCreator := &notificationConnectionCreatorImpl{
		storageURI:           settings.URI,
		minReconnectInterval: settings.Notification.MinReconnectInterval,
		maxReconnectInterval: settings.Notification.MaxReconnectInterval,
	}
	if err != nil {
		return nil, err
	}

	return &Notificator{
		queueSize:       settings.Notification.QueuesSize,
		connectionMutex: &sync.Mutex{},
		consumersMutex:  &sync.Mutex{},
		consumers: &consumers{
			queues:    make(map[string][]storage.NotificationQueue),
			platforms: make([]*types.Platform, 0),
		},
		storage:           ns,
		connectionCreator: connectionCreator,
		lastKnownRevision: types.INVALIDREVISION,
	}, nil
}

// Start starts the Notificator. It must not be called concurrently.
func (n *Notificator) Start(ctx context.Context, group *sync.WaitGroup) error {
	if n.ctx != nil {
		return errors.New("notificator already started")
	}
	n.ctx = ctx
	if err := n.openConnection(); err != nil {
		return fmt.Errorf("could not open connection to database %v", err)
	}
	startInWaitGroup(n.awaitTermination, group)
	return nil
}

func (n *Notificator) addConsumer(platform *types.Platform, queue storage.NotificationQueue) int64 {
	n.consumersMutex.Lock()
	defer n.consumersMutex.Unlock()
	n.consumers.Add(platform, queue)
	return atomic.LoadInt64(&n.lastKnownRevision)
}

func (n *Notificator) RegisterConsumer(platform *types.Platform, lastKnownRevision int64) (storage.NotificationQueue, int64, error) {
	if atomic.LoadInt32(&n.isConnected) == aFalse {
		return nil, types.INVALIDREVISION, errors.New("cannot register consumer - Notificator is not running")
	}
	queue, err := storage.NewNotificationQueue(n.queueSize)
	if err != nil {
		return nil, types.INVALIDREVISION, err
	}
	if err = n.startListening(); err != nil {
		return nil, types.INVALIDREVISION, fmt.Errorf("listen to %s channel failed %v", postgresChannel, err)
	}
	lastKnownSMRevision := n.addConsumer(platform, queue)
	defer func() {
		if err != nil {
			if errUnregisterConsumer := n.UnregisterConsumer(queue); errUnregisterConsumer != nil {
				log.C(n.ctx).WithError(errUnregisterConsumer).Errorf("Could not unregister notification consumer %s", queue.ID())
			}
		}
	}()
	if lastKnownRevision > lastKnownSMRevision {
		log.C(n.ctx).Debug("lastKnownRevision is grater than the one SM knows")
		return nil, types.INVALIDREVISION, util.ErrInvalidNotificationRevision
	}
	if lastKnownRevision == types.INVALIDREVISION {
		return queue, lastKnownSMRevision, nil
	}
	if _, err := n.storage.GetNotificationByRevision(n.ctx, lastKnownRevision); err != nil {
		if err == util.ErrNotFoundInStorage {
			log.C(n.ctx).WithError(err).Debugf("notification with revision %d not found in storage", lastKnownRevision)
			return nil, types.INVALIDREVISION, util.ErrInvalidNotificationRevision
		}
		return nil, types.INVALIDREVISION, err
	}

	missedNotifications, err := n.storage.ListNotifications(n.ctx, platform.ID, lastKnownRevision, lastKnownSMRevision)
	if err != nil {
		return nil, types.INVALIDREVISION, err
	}
	filteredMissedNotification := make([]*types.Notification, 0, len(missedNotifications))
	var recipients []*types.Platform
	for _, notification := range missedNotifications {
		recipients = []*types.Platform{platform}
		for _, filter := range n.notificationFilters {
			recipients = filter(recipients, notification)
			if len(recipients) == 0 {
				break
			}
		}
		if len(recipients) != 0 {
			filteredMissedNotification = append(filteredMissedNotification, notification)
		}
	}

	if n.queueSize < len(filteredMissedNotification) {
		log.C(n.ctx).Debugf("too many missed notifications %d", len(filteredMissedNotification))
		return nil, types.INVALIDREVISION, util.ErrInvalidNotificationRevision
	}

	queueWithMissedNotifications, err := storage.NewNotificationQueue(n.queueSize)
	if err != nil {
		return nil, types.INVALIDREVISION, err
	}
	for _, notification := range filteredMissedNotification {
		err = queueWithMissedNotifications.Enqueue(notification)
		if err != nil {
			return nil, types.INVALIDREVISION, err
		}
	}

	n.consumersMutex.Lock()
	defer n.consumersMutex.Unlock()
	for {
		select {
		case notification, ok := <-queue.Channel():
			if !ok {
				err = errors.New("")
				return nil, types.INVALIDREVISION, err
			}
			err = queueWithMissedNotifications.Enqueue(notification)
			if err != nil {
				return nil, types.INVALIDREVISION, err
			}
		default:
			err = n.consumers.ReplaceQueue(queue.ID(), queueWithMissedNotifications)
			if err != nil {
				return nil, types.INVALIDREVISION, err
			}
			return queueWithMissedNotifications, lastKnownSMRevision, nil
		}
	}
}

func (n *Notificator) UnregisterConsumer(queue storage.NotificationQueue) error {
	n.consumersMutex.Lock()
	defer n.consumersMutex.Unlock()
	n.consumers.Delete(queue)
	queue.Close()
	if n.consumers.Len() == 0 {
		return n.stopListening()
	}
	return nil
}

// RegisterFilter adds new notification filter. It must not be called concurrently.
func (n *Notificator) RegisterFilter(f storage.NotificationFilterFunc) {
	n.notificationFilters = append(n.notificationFilters, f)
}

func (n *Notificator) closeAllConsumers() {
	n.consumersMutex.Lock()
	defer n.consumersMutex.Unlock()

	platformConsumers := n.consumers.Clear()
	for _, platformConsumers := range platformConsumers {
		for _, queue := range platformConsumers {
			queue.Close()
		}
	}
}

func (n *Notificator) setConnection(conn notificationConnection.NotificationConnection) {
	n.connectionMutex.Lock()
	defer n.connectionMutex.Unlock()
	n.connection = conn
}

func (n *Notificator) openConnection() error {
	connection := n.connectionCreator.NewConnection(func(isConnected bool, err error) {
		if isConnected {
			atomic.StoreInt32(&n.isConnected, aTrue)
		} else {
			atomic.StoreInt32(&n.isConnected, aFalse)
			log.C(n.ctx).WithError(err).Info("connection to db closed, closing all consumers")
			n.closeAllConsumers()
		}
	})
	n.setConnection(connection)
	return nil
}

type notifyEventPayload struct {
	PlatformID     string `json:"platform_id"`
	NotificationID string `json:"notification_id"`
	Revision       int64  `json:"revision"`
}

func (n *Notificator) processNotifications(notificationChannel <-chan *pq.Notification) {
	defer func() {
		atomic.StoreInt32(&n.isListening, aFalse)
	}()
	for pqNotification := range notificationChannel {
		if pqNotification == nil {
			continue
		}
		payload, err := getPayload(pqNotification.Extra)
		if err != nil {
			log.C(n.ctx).WithError(err).Error("could not unmarshal notification payload")
			n.closeAllConsumers() // Ensures no notifications are lost
		} else {
			if err = n.processNotificationPayload(payload); err != nil {
				log.C(n.ctx).WithError(err).Error("closing consumers")
				n.closeAllConsumers() // Ensures no notifications are lost
			}
		}
	}
}

func getPayload(data string) (*notifyEventPayload, error) {
	payload := &notifyEventPayload{}
	if err := json.Unmarshal([]byte(data), payload); err != nil {
		return nil, err
	}
	return payload, nil
}

func (n *Notificator) processNotificationPayload(payload *notifyEventPayload) error {
	notificationPlatformID := payload.PlatformID
	notificationID := payload.NotificationID

	n.consumersMutex.Lock()
	defer n.consumersMutex.Unlock()
	atomic.StoreInt64(&n.lastKnownRevision, payload.Revision)

	recipients := n.getRecipients(notificationPlatformID)
	if len(recipients) == 0 {
		return nil
	}
	notification, err := n.storage.GetNotification(n.ctx, notificationID)
	if err != nil {
		return fmt.Errorf("notification %s could not be retrieved from the DB: %v", notificationID, err.Error())
	}
	for _, filter := range n.notificationFilters {
		recipients = filter(recipients, notification)
	}
	for _, platform := range recipients {
		n.sendNotificationToPlatformConsumers(n.consumers.GetQueuesForPlatform(platform.ID), notification)
	}
	return nil
}

func (n *Notificator) getRecipients(platformID string) []*types.Platform {
	if platformID == "" {
		return n.consumers.platforms
	}
	platform := n.consumers.GetPlatform(platformID)
	if platform == nil {
		return nil
	}
	return []*types.Platform{platform}
}

func (n *Notificator) sendNotificationToPlatformConsumers(platformConsumers []storage.NotificationQueue, notification *types.Notification) {
	for _, consumer := range platformConsumers {
		if err := consumer.Enqueue(notification); err != nil {
			log.C(n.ctx).WithError(err).Infof("consumer %s notification queue returned error %v", consumer.ID(), err)
			consumer.Close()
		}
	}
}

func startInWaitGroup(f func(), group *sync.WaitGroup) {
	group.Add(1)
	go func() {
		defer group.Done()
		f()
	}()
}

func (n *Notificator) awaitTermination() {
	<-n.ctx.Done()
	logger := log.C(n.ctx)
	logger.Info("context cancelled, stopping Notificator...")
	n.stopConnection()
}

func (n *Notificator) stopConnection() {
	err := n.stopListening()
	logger := log.C(n.ctx)
	if err != nil {
		logger.WithError(err).Info("could not unlisten notification channel")
	}
	n.connectionMutex.Lock()
	defer n.connectionMutex.Unlock()
	if err = n.connection.Close(); err != nil {
		logger.WithError(err).Info("could not close db connection")
	}
}

func (n *Notificator) stopListening() error {
	n.connectionMutex.Lock()
	defer n.connectionMutex.Unlock()
	if atomic.LoadInt32(&n.isListening) == aFalse {
		return nil
	}
	return n.connection.Unlisten(postgresChannel)
}

func (n *Notificator) startListening() error {
	n.connectionMutex.Lock()
	defer n.connectionMutex.Unlock()
	if atomic.LoadInt32(&n.isListening) == aTrue {
		return nil
	}
	err := n.connection.Listen(postgresChannel)
	if err != nil {
		return err
	}
	lastKnownRevision, err := n.storage.GetLastRevision(n.ctx)
	if err != nil {
		if errUnlisten := n.connection.Unlisten(postgresChannel); errUnlisten != nil {
			log.C(n.ctx).WithError(errUnlisten).Errorf("could not unlisten %s channel", postgresChannel)
		}
		return err
	}
	atomic.StoreInt64(&n.lastKnownRevision, lastKnownRevision)
	atomic.StoreInt32(&n.isListening, aTrue)
	go n.processNotifications(n.connection.NotificationChannel())
	return nil
}

type consumers struct {
	queues    map[string][]storage.NotificationQueue
	platforms []*types.Platform
}

func (c *consumers) find(queueID string) (string, int) {
	for platformID, notificationQueues := range c.queues {
		for index, queue := range notificationQueues {
			if queue.ID() == queueID {
				return platformID, index
			}
		}
	}
	return "", -1
}

func (c *consumers) ReplaceQueue(queueID string, newQueue storage.NotificationQueue) error {
	platformID, queueIndex := c.find(queueID)
	if queueIndex == -1 {
		return fmt.Errorf("could not find consumer with id %s", queueID)
	}
	c.queues[platformID][queueIndex] = newQueue
	return nil
}

func (c *consumers) Delete(queue storage.NotificationQueue) {
	platformIDToDelete, queueIndex := c.find(queue.ID())
	if queueIndex == -1 {
		return
	}
	platformConsumers := c.queues[platformIDToDelete]
	c.queues[platformIDToDelete] = append(platformConsumers[:queueIndex], platformConsumers[queueIndex+1:]...)

	if len(c.queues[platformIDToDelete]) == 0 {
		delete(c.queues, platformIDToDelete)
		for index, platform := range c.platforms {
			if platform.ID == platformIDToDelete {
				c.platforms = append(c.platforms[:index], c.platforms[index+1:]...)
				break
			}
		}
	}
}

func (c *consumers) Add(platform *types.Platform, queue storage.NotificationQueue) {
	if len(c.queues[platform.ID]) == 0 {
		c.platforms = append(c.platforms, platform)
	}
	c.queues[platform.ID] = append(c.queues[platform.ID], queue)
}

func (c *consumers) Clear() map[string][]storage.NotificationQueue {
	allQueues := c.queues
	c.queues = make(map[string][]storage.NotificationQueue)
	c.platforms = make([]*types.Platform, 0)
	return allQueues
}

func (c *consumers) Len() int {
	return len(c.queues)
}

func (c *consumers) GetPlatform(platformID string) *types.Platform {
	for _, platform := range c.platforms {
		if platform.ID == platformID {
			return platform
		}
	}
	return nil
}

func (c *consumers) GetQueuesForPlatform(platformID string) []storage.NotificationQueue {
	return c.queues[platformID]
}
