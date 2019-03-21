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

package filters

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Peripli/service-manager/pkg/query"

	"github.com/tidwall/gjson"

	"github.com/Peripli/service-manager/pkg/util"

	"github.com/Peripli/service-manager/pkg/types"
	"github.com/gofrs/uuid"

	"github.com/Peripli/service-manager/pkg/log"

	"github.com/Peripli/service-manager/storage"

	"github.com/Peripli/service-manager/pkg/web"
)

// PublicServicePlansFilter reconciles the state of the free plans offered by all service brokers registered in SM. The
// filter makes sure that a public visibility exists for each free plan present in SM DB.
type PublicServicePlansFilter struct {
	Repository              storage.Repository
	IsCatalogPlanPublicFunc func(broker *types.ServiceBroker, catalogService *types.ServiceOffering, catalogPlan *types.ServicePlan) (bool, error)
}

func (pspf *PublicServicePlansFilter) Name() string {
	return "PublicServicePlansFilter"
}

func (pspf *PublicServicePlansFilter) Run(req *web.Request, next web.Handler) (*web.Response, error) {
	response, err := next.Handle(req)
	if err != nil {
		return nil, err
	}
	ctx := req.Context()
	brokerID := gjson.GetBytes(response.Body, "id").String()
	log.C(ctx).Debugf("Reconciling public plans for broker with id: %s", brokerID)
	if err := pspf.Repository.InTransaction(ctx, func(ctx context.Context, storage storage.Warehouse) error {
		soRepository := storage.ServiceOffering()

		broker, err := pspf.Repository.Get(ctx, types.ServiceBrokerType, brokerID)
		if err != nil {
			return util.HandleStorageError(err, "broker")
		}
		catalog, err := soRepository.ListWithServicePlansByBrokerID(ctx, brokerID)
		if err != nil {
			return err
		}
		for _, serviceOffering := range catalog {
			for _, servicePlan := range serviceOffering.Plans {
				planID := servicePlan.ID
				isPublic, err := pspf.IsCatalogPlanPublicFunc(broker.(*types.ServiceBroker), serviceOffering, servicePlan)
				if err != nil {
					return err
				}

				hasPublicVisibility := false
				byServicePlanID := query.ByField(query.EqualsOperator, "service_plan_id", planID)
				visibilitiesForPlan, err := storage.List(ctx, types.VisibilityType, byServicePlanID)
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
							if _, err := storage.Delete(ctx, types.VisibilityType, byVisibilityID); err != nil {
								return err
							}
						}
					} else {
						if visibility.PlatformID == "" {
							if _, err := storage.Delete(ctx, types.VisibilityType, byVisibilityID); err != nil {
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
					planID, err := storage.Create(ctx, &types.Visibility{
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

					log.C(ctx).Debugf("Created new public visibility for broker with id %s and plan with id %s", brokerID, planID)
				}
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}
	log.C(ctx).Debugf("Successfully finished reconciling public plans for broker with id %s", brokerID)
	return response, nil
}

func (pspf *PublicServicePlansFilter) FilterMatchers() []web.FilterMatcher {
	return []web.FilterMatcher{
		{
			Matchers: []web.Matcher{
				web.Path(web.BrokersURL + "/**"),
				web.Methods(http.MethodPost, http.MethodPatch),
			},
		},
	}
}