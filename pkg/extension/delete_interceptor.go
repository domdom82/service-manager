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

package extension

import (
	"context"
	"fmt"

	"github.com/Peripli/service-manager/pkg/types"

	"github.com/Peripli/service-manager/pkg/query"

	"github.com/Peripli/service-manager/storage"
)

type namedDeleteAPIFunc struct {
	Name string
	Func func(InterceptDeleteOnAPI) InterceptDeleteOnAPI
}

type namedDeleteTxFunc struct {
	Name string
	Func func(InterceptDeleteOnTransaction) InterceptDeleteOnTransaction
}

type deleteHookOnAPIHandler struct {
	DeleteHookOnAPIFuncs []*namedDeleteAPIFunc
	DeleteHookOnTxFuncs  []*namedDeleteTxFunc
}

func (c *deleteHookOnAPIHandler) OnAPIDelete(f InterceptDeleteOnAPI) InterceptDeleteOnAPI {
	for i := range c.DeleteHookOnAPIFuncs {
		f = c.DeleteHookOnAPIFuncs[len(c.DeleteHookOnAPIFuncs)-1-i].Func(f)
	}
	return f
}

func (c *deleteHookOnAPIHandler) OnTransactionDelete(f InterceptDeleteOnTransaction) InterceptDeleteOnTransaction {
	for i := range c.DeleteHookOnTxFuncs {
		f = c.DeleteHookOnTxFuncs[len(c.DeleteHookOnTxFuncs)-1-i].Func(f)
	}
	return f
}

// UnionDeleteInterceptor returns a function which spawns all delete interceptors, sorts them and wraps them into one.
func UnionDeleteInterceptor(providers []DeleteInterceptorProvider) func() DeleteInterceptor {
	return func() DeleteInterceptor {
		c := &deleteHookOnAPIHandler{}
		c.DeleteHookOnAPIFuncs = make([]*namedDeleteAPIFunc, 0, len(providers))
		c.DeleteHookOnTxFuncs = make([]*namedDeleteTxFunc, 0, len(providers))

		for _, p := range providers {
			hook := p.Provide()
			positionAPIType := PositionNone
			positionTxType := PositionNone
			nameAPI := ""
			nameTx := ""

			if orderedProvider, isOrdered := p.(Ordered); isOrdered {
				positionAPIType, nameAPI = orderedProvider.PositionAPI()
				positionTxType, nameTx = orderedProvider.PositionTransaction()
			}

			c.insertAPIFunc(positionAPIType, nameAPI, &namedDeleteAPIFunc{
				Name: p.Name(),
				Func: hook.OnAPIDelete,
			})
			c.insertTxFunc(positionTxType, nameTx, &namedDeleteTxFunc{
				Name: p.Name(),
				Func: hook.OnTransactionDelete,
			})
		}
		return c
	}
}

// DeleteInterceptorProvider provides DeleteInterceptors for each request
//go:generate counterfeiter . DeleteInterceptorProvider
type DeleteInterceptorProvider interface {
	Named
	Provide() DeleteInterceptor
}

// InterceptDeleteOnAPI hook for entity deletion outside of transaction
type InterceptDeleteOnAPI func(ctx context.Context, deletionCriteria ...query.Criterion) (types.ObjectList, error)

// InterceptDeleteOnTransaction hook for entity deletion in transaction
type InterceptDeleteOnTransaction func(ctx context.Context, txStorage storage.Warehouse, deletionCriteria ...query.Criterion) (types.ObjectList, error)

// DeleteInterceptor provides hooks on entity deletion
//go:generate counterfeiter . DeleteInterceptor
type DeleteInterceptor interface {
	OnAPIDelete(h InterceptDeleteOnAPI) InterceptDeleteOnAPI
	OnTransactionDelete(f InterceptDeleteOnTransaction) InterceptDeleteOnTransaction
}

func (c *deleteHookOnAPIHandler) insertAPIFunc(positionType PositionType, name string, h *namedDeleteAPIFunc) {
	if positionType == PositionNone {
		c.DeleteHookOnAPIFuncs = append(c.DeleteHookOnAPIFuncs, h)
		return
	}
	pos := c.findAPIFuncPosition(c.DeleteHookOnAPIFuncs, name)
	if pos == -1 {
		// TODO: Must validate on bootstrap
		panic(fmt.Errorf("could not find delete API hook with name %s", name))
	}
	c.DeleteHookOnAPIFuncs = append(c.DeleteHookOnAPIFuncs, nil)
	if positionType == PositionAfter {
		pos = pos + 1
	}
	copy(c.DeleteHookOnAPIFuncs[pos+1:], c.DeleteHookOnAPIFuncs[pos:])
	c.DeleteHookOnAPIFuncs[pos] = h
}

func (c *deleteHookOnAPIHandler) insertTxFunc(positionType PositionType, name string, h *namedDeleteTxFunc) {
	if positionType == PositionNone {
		c.DeleteHookOnTxFuncs = append(c.DeleteHookOnTxFuncs, h)
		return
	}
	pos := c.findTxFuncPosition(c.DeleteHookOnTxFuncs, name)
	if pos == -1 {
		// TODO: Must validate on bootstrap
		panic(fmt.Errorf("could not find delete transaction hook with name %s", name))
	}
	c.DeleteHookOnTxFuncs = append(c.DeleteHookOnTxFuncs, nil)
	if positionType == PositionAfter {
		pos = pos + 1
	}
	copy(c.DeleteHookOnTxFuncs[pos+1:], c.DeleteHookOnTxFuncs[pos:])
	c.DeleteHookOnTxFuncs[pos] = h
}

func (c *deleteHookOnAPIHandler) findAPIFuncPosition(funcs []*namedDeleteAPIFunc, name string) int {
	for i, f := range funcs {
		if f.Name == name {
			return i
		}
	}

	return -1
}

func (c *deleteHookOnAPIHandler) findTxFuncPosition(funcs []*namedDeleteTxFunc, name string) int {
	for i, f := range funcs {
		if f.Name == name {
			return i
		}
	}

	return -1
}
