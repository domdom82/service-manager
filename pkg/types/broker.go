/*
 * Copyright 2018 The Service Manager Authors
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

// Package types contains the Service Manager web entities
package types

import (
	"time"

	"errors"
)

// TODO this gogen (after applying the other TODOs) should be sufficient to generate a controller, a postgres.Entity and the whole storage layer
// TODO last parameter of the gogen should not be needed - if whoever defined the struct put Labels in it, then it supports labels.
// TODO you could also move out the fields that implement Object to a struct and have that embedded everywhere together with the methods
//go:generate smgen api broker labels
// Broker broker struct
type Broker struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	BrokerURL   string       `json:"broker_url"`
	Credentials *Credentials `json:"credentials,omitempty" structs:"-"`

	Services []*ServiceOffering `json:"services,omitempty" structs:"-"`

	Labels Labels `json:"labels,omitempty"`
}

func (e *Broker) GetUpdatedAt() time.Time {
	return e.UpdatedAt
}

func (e *Broker) GetCreatedAt() time.Time {
	return e.CreatedAt
}

func (e *Broker) SetID(id string) {
	e.ID = id
}

func (e *Broker) GetID() string {
	return e.ID
}

func (e *Broker) SetCreatedAt(time time.Time) {
	e.CreatedAt = time
}

func (e *Broker) SetUpdatedAt(time time.Time) {
	e.UpdatedAt = time
}

func (e *Broker) SetCredentials(credentials *Credentials) {
	e.Credentials = credentials
}

// Validate implements InputValidator and verifies all mandatory fields are populated
func (b *Broker) Validate() error {
	if b.Name == "" {
		return errors.New("missing broker name")
	}
	if b.BrokerURL == "" {
		return errors.New("missing broker url")
	}

	if err := b.Labels.Validate(); err != nil {
		return err
	}

	if b.Credentials == nil {
		return errors.New("missing credentials")
	}
	return b.Credentials.Validate()
}
