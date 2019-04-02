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

package web_test

import (
	"github.com/Peripli/service-manager/api/base"
	"github.com/Peripli/service-manager/pkg/extension/extensionfakes"
	"github.com/Peripli/service-manager/pkg/types"
	"github.com/Peripli/service-manager/pkg/web"
	"github.com/Peripli/service-manager/storage/storagefakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("API", func() {
	var (
		api *web.API
	)

	BeforeEach(func() {
		api = &web.API{}
	})

	filterNames := func() []string {
		var names []string
		for i := range api.Filters {
			names = append(names, api.Filters[i].Name())
		}
		return names
	}

	Describe("RegisterControllers", func() {
		It("increases controllers count", func() {
			originalCount := len(api.Controllers)
			api.RegisterControllers(&testController{})
			Expect(len(api.Controllers)).To(Equal(originalCount + 1))
		})
	})

	Describe("RegisterPlugins", func() {
		Context("When argument is empty plugin", func() {
			It("Panics", func() {
				Expect(func() {
					api.RegisterPlugins(&invalidPlugin{})
				}).To(Panic())
			})
		})

		Context("When plugin is valid", func() {
			It("Successfully registers plugin", func() {
				originalCount := len(api.Filters)
				api.RegisterPlugins(&validPlugin{})
				Expect(len(api.Filters)).To(Equal(originalCount + 8))
			})
		})

		Context("When plugin with the same name is already registered", func() {
			It("Panics", func() {
				registerPlugin := func() {
					api.RegisterPlugins(&validPlugin{})
				}
				registerPlugin()
				Expect(registerPlugin).To(Panic())
			})
		})
	})

	Describe("Replace Filter", func() {
		Context("When filter with such name does not exist", func() {
			It("Panics", func() {
				replaceFilter := func() {
					api.ReplaceFilter("some-filter", &testFilter{"testFilter"})
				}
				Expect(replaceFilter).To(Panic())
			})
		})
		Context("When filter with such name exists", func() {
			It("Replaces the filter", func() {
				filter := &testFilter{"testFilter"}
				newFilter := &testFilter{"testFilter2"}
				api.RegisterFilters(filter)
				api.ReplaceFilter(filter.Name(), newFilter)
				names := filterNames()
				Expect(names).To(ConsistOf([]string{newFilter.Name()}))
			})
		})
	})

	Describe("Register Filter Before", func() {
		Context("When filter with such name does not exist", func() {
			It("Panics", func() {
				replaceFilter := func() {
					api.RegisterFiltersBefore("some-filter", &testFilter{"testFilter"})
				}
				Expect(replaceFilter).To(Panic())
			})
		})

		Context("When filter with such name exists", func() {
			It("Adds a filter before it", func() {
				filter1 := &testFilter{"testFilter1"}
				filter2 := &testFilter{"testFilter2"}
				filter3 := &testFilter{"testFilter3"}
				api.RegisterFilters(filter1, filter2)
				api.RegisterFiltersBefore(filter2.Name(), filter3)
				Expect(filterNames()).To(Equal([]string{filter1.Name(), filter3.Name(), filter2.Name()}))
			})

			It("Adds multiple filters before it", func() {
				filter1 := &testFilter{"testFilter1"}
				filter2 := &testFilter{"testFilter2"}
				filter3 := &testFilter{"testFilter3"}
				api.RegisterFilters(filter1)
				api.RegisterFiltersBefore(filter1.Name(), filter2, filter3)
				Expect(filterNames()).To(Equal([]string{filter2.Name(), filter3.Name(), filter1.Name()}))
			})
		})
	})

	Describe("Register Filter After", func() {
		Context("When filter with such name does not exist", func() {
			It("Panics", func() {
				replaceFilter := func() {
					api.RegisterFiltersAfter("some-filter", &testFilter{"testFilter"})
				}
				Expect(replaceFilter).To(Panic())
			})
		})
		Context("When filter with such name exists", func() {
			It("Adds a filter after it", func() {
				filter1 := &testFilter{"testFilter1"}
				filter2 := &testFilter{"testFilter2"}
				filter3 := &testFilter{"testFilter3"}
				api.RegisterFilters(filter1, filter2)
				api.RegisterFiltersAfter(filter1.Name(), filter3)
				Expect(filterNames()).To(Equal([]string{filter1.Name(), filter3.Name(), filter2.Name()}))
			})

			It("Adds multiple filters after it", func() {
				filter1 := &testFilter{"testFilter1"}
				filter2 := &testFilter{"testFilter2"}
				filter3 := &testFilter{"testFilter3"}
				filter4 := &testFilter{"testFilter4"}
				api.RegisterFilters(filter1, filter2)
				api.RegisterFiltersAfter(filter1.Name(), filter3, filter4)
				Expect(filterNames()).To(Equal([]string{filter1.Name(), filter3.Name(), filter4.Name(), filter2.Name()}))
			})
		})
	})

	Describe("Remove Filter", func() {
		Context("When filter with such name doest not exist", func() {
			It("Panics", func() {
				removeFilter := func() {
					api.RemoveFilter("some-filter")
				}
				Expect(removeFilter).To(Panic())
			})
		})
		Context("When filter exists", func() {
			It("Should remove it", func() {
				filter := &testFilter{"testFilter"}
				api.RegisterFilters(filter)
				names := filterNames()
				Expect(names).To(ConsistOf(filter.Name()))

				api.RemoveFilter(filter.Name())
				names = filterNames()
				Expect(names).To(BeEmpty())
			})
		})
	})

	Describe("RegisterFilters", func() {
		Context("When filter with such name does not exist", func() {
			It("increases filter count if successful", func() {
				originalCount := len(api.Filters)
				api.RegisterFilters(&testFilter{"testFilter"})
				Expect(len(api.Filters)).To(Equal(originalCount + 1))
			})
		})

		Context("When filter with such name already exists", func() {
			It("Panics", func() {
				registerFilter := func() {
					api.RegisterFilters(&testFilter{"testFilter"})
				}
				registerFilter()
				Expect(registerFilter).To(Panic())
			})
		})

		Context("When filter name contains :", func() {
			It("Panics", func() {
				registerFilter := func() {
					api.RegisterFilters(&testFilter{"name:"})
				}
				Expect(registerFilter).To(Panic())
			})
		})

		Context("When filter name is empty", func() {
			It("Panics", func() {
				registerFilter := func() {
					api.RegisterFilters(&testFilter{""})
				}
				Expect(registerFilter).To(Panic())
			})
		})
	})

	Describe("Register interceptor", func() {

		BeforeEach(func() {
			api.RegisterControllers(base.NewController(&storagefakes.FakeStorage{}, web.ServiceBrokersURL, func() types.Object {
				return &types.ServiceBroker{}
			}))
		})

		Context("Create interceptor", func() {
			Context("When provider with the same name is already registered", func() {
				It("Panics", func() {
					interceptor := &extensionfakes.FakeCreateInterceptorProvider{NameStub: func() string { return "Create" }}
					f := func() { api.RegisterCreateInterceptorProvider(types.ServiceBrokerType, interceptor).Apply() }
					f()
					Expect(f).To(Panic())
				})
			})
			Context("When registering before non-existing one", func() {
				It("Panics", func() {
					interceptor := &extensionfakes.FakeCreateInterceptorProvider{NameStub: func() string { return "Create" }}
					f := func() {
						api.RegisterCreateInterceptorProvider(types.ServiceBrokerType, interceptor).Before("no_such_interceptor").Apply()
					}
					Expect(f).To(Panic())
				})
			})
			Context("When registering on non-existing interceptable type", func() {
				It("Panics", func() {
					interceptor := &extensionfakes.FakeCreateInterceptorProvider{NameStub: func() string { return "Create" }}
					f := func() { api.RegisterCreateInterceptorProvider(types.PlatformType, interceptor).Apply() }
					Expect(f).To(Panic())
				})
			})
		})

		Context("Update interceptor", func() {
			Context("When provider with the same name is already registered", func() {
				It("Panics", func() {
					interceptor := &extensionfakes.FakeUpdateInterceptorProvider{NameStub: func() string { return "Update" }}
					f := func() { api.RegisterUpdateInterceptorProvider(types.ServiceBrokerType, interceptor).Apply() }
					f()
					Expect(f).To(Panic())
				})
			})
			Context("When registering before non-existing one", func() {
				It("Panics", func() {
					interceptor := &extensionfakes.FakeUpdateInterceptorProvider{NameStub: func() string { return "Update" }}
					f := func() {
						api.RegisterUpdateInterceptorProvider(types.ServiceBrokerType, interceptor).Before("no_such_interceptor").Apply()
					}
					Expect(f).To(Panic())
				})
			})

			Context("When registering on non-existing interceptable type", func() {
				It("Panics", func() {
					interceptor := &extensionfakes.FakeUpdateInterceptorProvider{NameStub: func() string { return "Update" }}
					f := func() { api.RegisterUpdateInterceptorProvider(types.PlatformType, interceptor).Apply() }
					Expect(f).To(Panic())
				})
			})
		})

		Context("Update interceptor", func() {
			Context("When provider with the same name is already registered", func() {
				It("Panics", func() {
					interceptor := &extensionfakes.FakeDeleteInterceptorProvider{NameStub: func() string { return "Delete" }}
					f := func() { api.RegisterDeleteInterceptorProvider(types.ServiceBrokerType, interceptor).Apply() }
					f()
					Expect(f).To(Panic())
				})
			})
			Context("When registering before non-existing one", func() {
				It("Panics", func() {
					interceptor := &extensionfakes.FakeDeleteInterceptorProvider{NameStub: func() string { return "Delete" }}
					f := func() {
						api.RegisterDeleteInterceptorProvider(types.ServiceBrokerType, interceptor).Before("no_such_interceptor").Apply()
					}
					Expect(f).To(Panic())
				})
			})
			Context("When registering on non-existing interceptable type", func() {
				It("Panics", func() {
					interceptor := &extensionfakes.FakeDeleteInterceptorProvider{NameStub: func() string { return "Create" }}
					f := func() { api.RegisterDeleteInterceptorProvider(types.PlatformType, interceptor).Apply() }
					Expect(f).To(Panic())
				})
			})
		})
	})
})

type testController struct {
}

func (c *testController) Routes() []web.Route {
	return []web.Route{}
}

type testFilter struct {
	name string
}

func (tf testFilter) Name() string {
	return tf.name
}

func (tf testFilter) Run(request *web.Request, next web.Handler) (*web.Response, error) {
	return next.Handle(request)
}

func (tf testFilter) FilterMatchers() []web.FilterMatcher {
	return []web.FilterMatcher{}
}

type invalidPlugin struct {
}

func (p *invalidPlugin) Name() string {
	return "invalidPlugin"
}

type validPlugin struct {
}

func (c *validPlugin) UpdateService(request *web.Request, next web.Handler) (*web.Response, error) {
	return next.Handle(request)
}

func (c *validPlugin) Unbind(request *web.Request, next web.Handler) (*web.Response, error) {
	return next.Handle(request)
}

func (c *validPlugin) Bind(request *web.Request, next web.Handler) (*web.Response, error) {
	return next.Handle(request)
}

func (c *validPlugin) FetchBinding(request *web.Request, next web.Handler) (*web.Response, error) {
	return next.Handle(request)
}

func (c *validPlugin) Deprovision(request *web.Request, next web.Handler) (*web.Response, error) {
	return next.Handle(request)
}

func (c *validPlugin) Provision(request *web.Request, next web.Handler) (*web.Response, error) {
	return next.Handle(request)
}

func (c *validPlugin) FetchService(request *web.Request, next web.Handler) (*web.Response, error) {
	return next.Handle(request)
}

func (c *validPlugin) FetchCatalog(request *web.Request, next web.Handler) (*web.Response, error) {
	return next.Handle(request)
}

func (c *validPlugin) Name() string {
	return "validPlugin"
}
