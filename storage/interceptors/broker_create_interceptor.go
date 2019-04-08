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

package interceptors

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Peripli/service-manager/pkg/types"
	"github.com/Peripli/service-manager/pkg/util"
	"github.com/Peripli/service-manager/storage"
	"github.com/gofrs/uuid"
	osbc "github.com/pmorie/go-open-service-broker-client/v2"
)

const CreateBrokerInterceptorName = "create-broker"

type BrokerCreateInterceptorProvider struct {
	OsbClientCreateFunc osbc.CreateFunc
}

func (c *BrokerCreateInterceptorProvider) Provide() storage.CreateInterceptor {
	return &CreateBrokerInterceptor{
		OSBClientCreateFunc: c.OsbClientCreateFunc,
	}
}

type CreateBrokerInterceptor struct {
	OSBClientCreateFunc osbc.CreateFunc
	serviceOfferings    []*types.ServiceOffering
}

func (c *CreateBrokerInterceptor) Name() string {
	return CreateBrokerInterceptorName
}

func (c *CreateBrokerInterceptor) AroundTxCreate(h storage.InterceptCreateAroundTxFunc) storage.InterceptCreateAroundTxFunc {
	return func(ctx context.Context, obj types.Object) (types.Object, error) {
		broker := obj.(*types.ServiceBroker)
		catalog, err := getBrokerCatalog(ctx, c.OSBClientCreateFunc, broker) // keep catalog to be stored later
		if err != nil {
			return nil, err
		}
		if c.serviceOfferings, err = osbCatalogToOfferings(catalog, broker); err != nil {
			return nil, err
		}
		broker.Services = c.serviceOfferings
		return h(ctx, broker)
	}
}

func osbCatalogToOfferings(catalog *osbc.CatalogResponse, broker types.Object) ([]*types.ServiceOffering, error) {
	var result []*types.ServiceOffering
	for serviceIndex := range catalog.Services {
		service := catalog.Services[serviceIndex]
		serviceOffering := &types.ServiceOffering{}
		err := osbcCatalogServiceToServiceOffering(serviceOffering, &service)
		if err != nil {
			return nil, err
		}
		serviceUUID, err := uuid.NewV4()
		if err != nil {
			return nil, fmt.Errorf("could not generate GUID for service: %s", err)
		}
		serviceOffering.ID = serviceUUID.String()
		serviceOffering.CreatedAt = broker.GetCreatedAt()
		serviceOffering.UpdatedAt = broker.GetUpdatedAt()
		serviceOffering.BrokerID = broker.GetID()

		if err := serviceOffering.Validate(); err != nil {
			return nil, &util.HTTPError{
				ErrorType:   "BadRequest",
				Description: fmt.Sprintf("service offering constructed during catalog insertion for broker %s is invalid: %s", broker.GetID(), err),
				StatusCode:  http.StatusBadRequest,
			}
		}

		for planIndex := range service.Plans {
			servicePlan := &types.ServicePlan{}
			err := osbcCatalogPlanToServicePlan(servicePlan, &catalogPlanWithServiceOfferingID{
				Plan:            &service.Plans[planIndex],
				ServiceOffering: serviceOffering,
			})
			if err != nil {
				return nil, err
			}
			planUUID, err := uuid.NewV4()
			if err != nil {
				return nil, fmt.Errorf("could not generate GUID for service_plan: %s", err)
			}
			servicePlan.ID = planUUID.String()
			servicePlan.CreatedAt = broker.GetCreatedAt()
			servicePlan.UpdatedAt = broker.GetUpdatedAt()

			if err := servicePlan.Validate(); err != nil {
				return nil, &util.HTTPError{
					ErrorType:   "BadRequest",
					Description: fmt.Sprintf("service plan constructed during catalog insertion for broker %s is invalid: %s", broker.GetID(), err),
					StatusCode:  http.StatusBadRequest,
				}
			}
			serviceOffering.Plans = append(serviceOffering.Plans, servicePlan)
		}
		result = append(result, serviceOffering)
	}
	return result, nil
}

func (c *CreateBrokerInterceptor) OnTxCreate(f storage.InterceptCreateOnTxFunc) storage.InterceptCreateOnTxFunc {
	return func(ctx context.Context, storage storage.Repository, broker types.Object) error {
		if err := f(ctx, storage, broker); err != nil {
			return err
		}
		for serviceIndex := range c.serviceOfferings {
			service := c.serviceOfferings[serviceIndex]
			if _, err := storage.Create(ctx, service); err != nil {
				return util.HandleStorageError(err, "service_offering")
			}
			for planIndex := range service.Plans {
				servicePlan := service.Plans[planIndex]
				if _, err := storage.Create(ctx, servicePlan); err != nil {
					return util.HandleStorageError(err, "service_plan")
				}
			}
		}
		return nil
	}
}
