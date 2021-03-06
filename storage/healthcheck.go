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

package storage

import "github.com/Peripli/service-manager/pkg/health"

// HealthIndicator returns a new indicator for the storage
type HealthIndicator struct {
	Pinger Pinger
}

// Name returns the name of the storage component
func (i *HealthIndicator) Name() string {
	return "storage"
}

// Health returns the health of the storage component
func (i *HealthIndicator) Health() *health.Health {
	err := i.Pinger.Ping()
	healthz := health.New()
	if err != nil {
		return healthz.WithError(err).WithDetail("message", "TransactionalRepository ping failed")
	}
	return healthz.Up()
}
