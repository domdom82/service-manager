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
	"time"

	"github.com/Peripli/service-manager/pkg/extension"
	"github.com/Peripli/service-manager/pkg/log"
	"github.com/Peripli/service-manager/pkg/query"
	"github.com/Peripli/service-manager/pkg/types"
	"github.com/Peripli/service-manager/pkg/util"
	"github.com/Peripli/service-manager/storage"
	"github.com/gofrs/uuid"
)

const (
	PublicPlanCreateInterceptorProviderName = "public-plan-create"
	PublicPlanUpdateInterceptorProviderName = "public-plan-update"
)

type publicPlanProcessor func(broker *types.ServiceBroker, catalogService *types.ServiceOffering, catalogPlan *types.ServicePlan) (bool, error)

type PublicPlanCreateInterceptorProvider struct {
	IsCatalogPlanPublicFunc publicPlanProcessor
}

func (p *PublicPlanCreateInterceptorProvider) Name() string {
	return PublicPlanCreateInterceptorProviderName
}

func (p *PublicPlanCreateInterceptorProvider) Provide() extension.CreateInterceptor {
	return &publicPlanCreateInterceptor{
		isCatalogPlanPublicFunc: p.IsCatalogPlanPublicFunc,
	}
}

type PublicPlanUpdateInterceptorProvider struct {
	IsCatalogPlanPublicFunc publicPlanProcessor
}

func (p *PublicPlanUpdateInterceptorProvider) Name() string {
	return PublicPlanUpdateInterceptorProviderName
}

func (p *PublicPlanUpdateInterceptorProvider) Provide() extension.UpdateInterceptor {
	return &publicPlanUpdateInterceptor{
		isCatalogPlanPublicFunc: p.IsCatalogPlanPublicFunc,
	}
}

type publicPlanCreateInterceptor struct {
	isCatalogPlanPublicFunc publicPlanProcessor
}

func (p *publicPlanCreateInterceptor) OnAPICreate(h extension.InterceptCreateOnAPI) extension.InterceptCreateOnAPI {
	return h
}

func (p *publicPlanCreateInterceptor) OnTxCreate(f extension.InterceptCreateOnTx) extension.InterceptCreateOnTx {
	return func(ctx context.Context, txStorage storage.Warehouse, newObject types.Object) error {
		if err := f(ctx, txStorage, newObject); err != nil {
			return err
		}
		return resync(ctx, newObject.(*types.ServiceBroker), txStorage, p.isCatalogPlanPublicFunc)
	}
}

type publicPlanUpdateInterceptor struct {
	isCatalogPlanPublicFunc publicPlanProcessor
}

func (p *publicPlanUpdateInterceptor) OnAPIUpdate(h extension.InterceptUpdateOnAPI) extension.InterceptUpdateOnAPI {
	return h
}

func (p *publicPlanUpdateInterceptor) OnTxUpdate(f extension.InterceptUpdateOnTx) extension.InterceptUpdateOnTx {
	return func(ctx context.Context, txStorage storage.Warehouse, oldObject types.Object, changes *extension.UpdateContext) (types.Object, error) {
		result, err := f(ctx, txStorage, oldObject, changes)
		if err != nil {
			return nil, err
		}
		return result, resync(ctx, result.(*types.ServiceBroker), txStorage, p.isCatalogPlanPublicFunc)
	}
}

func resync(ctx context.Context, broker *types.ServiceBroker, txStorage storage.Warehouse, isCatalogPlanPublicFunc publicPlanProcessor) error {
	for _, serviceOffering := range broker.Services {
		for _, servicePlan := range serviceOffering.Plans {
			planID := servicePlan.ID
			isPublic, err := isCatalogPlanPublicFunc(broker, serviceOffering, servicePlan)
			if err != nil {
				return err
			}

			hasPublicVisibility := false
			byServicePlanID := query.ByField(query.EqualsOperator, "service_plan_id", planID)
			visibilitiesForPlan, err := txStorage.List(ctx, types.VisibilityType, byServicePlanID)
			if err != nil {
				return err
			}
			for i := 0; i < visibilitiesForPlan.Len(); i++ {
				visibility := visibilitiesForPlan.ItemAt(i).(*types.Visibility)
				byVisibilityID := query.ByField(query.EqualsOperator, "id", visibility.ID)
				if isPublic {
					if visibility.PlatformID == "" {
						hasPublicVisibility = true
						continue
					} else {
						if _, err := txStorage.Delete(ctx, types.VisibilityType, byVisibilityID); err != nil {
							return err
						}
					}
				} else {
					if visibility.PlatformID == "" {
						if _, err := txStorage.Delete(ctx, types.VisibilityType, byVisibilityID); err != nil {
							return err
						}
					} else {
						continue
					}
				}
			}

			if isPublic && !hasPublicVisibility {
				UUID, err := uuid.NewV4()
				if err != nil {
					return fmt.Errorf("could not generate GUID for visibility: %s", err)
				}

				currentTime := time.Now().UTC()
				planID, err := txStorage.Create(ctx, &types.Visibility{
					Base: types.Base{
						ID:        UUID.String(),
						UpdatedAt: currentTime,
						CreatedAt: currentTime,
					},
					ServicePlanID: servicePlan.ID,
				})
				if err != nil {
					return util.HandleStorageError(err, "visibility")
				}

				log.C(ctx).Debugf("Created new public visibility for broker with id %s and plan with id %s", broker.ID, planID)
			}
		}
	}
	return nil
}
