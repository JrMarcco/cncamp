package framework

import (
	"errors"
	"sync"
)

type serviceProviders map[string]ServiceProvider
type serviceInstances map[string]any

type Container interface {
	Bind(ServiceProvider) error
	IsBind(string) bool

	Make(string) (any, error)
	MustMake(string) any
	MakeNew(string, []any) (any, error)
}

type ServiceContainer struct {
	Container

	providers serviceProviders
	instances serviceInstances

	rwMutex sync.RWMutex
}

var _ Container = new(ServiceContainer)

func NewServiceContainer() *ServiceContainer {
	return &ServiceContainer{}
}

func (sc *ServiceContainer) Bind(sp ServiceProvider) error {
	sc.rwMutex.Lock()
	defer sc.rwMutex.Unlock()

	spName := sp.Name()
	sc.providers[spName] = sp

	if !sp.IsDefer() {
		if err := sp.Boot(sc); err != nil {
			return err
		}

		params := sp.Params(sc)
		method := sp.Register(sc)
		if instance, err := method(params); err != nil {
			return err
		} else {
			sc.instances[spName] = instance
		}
	}
	return nil
}

func (sc *ServiceContainer) IsBind(key string) bool {
	return sc.findServiceProvider(key) != nil
}

func (sc *ServiceContainer) Make(key string) (any, error) {
	return sc.make(key, nil, false)
}

func (sc *ServiceContainer) MustMake(key string) any {
	si, err := sc.make(key, nil, false)
	if err != nil {
		panic(err)
	}
	return si
}

func (sc *ServiceContainer) MakeNew(key string, params []any) (any, error) {
	return sc.make(key, params, true)
}

func (sc *ServiceContainer) make(key string, params []any, forceNew bool) (any, error) {
	sc.rwMutex.RLock()
	defer sc.rwMutex.RUnlock()

	sp := sc.findServiceProvider(key)
	if sp == nil {
		return nil, errors.New("contract " + key + " have not register")
	}

	if forceNew {
		return sc.newInstance(sp, params)
	}

	if inst, ok := sc.instances[key]; ok {
		return inst, nil
	}

	inst, err := sc.newInstance(sp, nil)
	if err != nil {
		return nil, err
	}
	return inst, nil
}

func (sc *ServiceContainer) findServiceProvider(key string) ServiceProvider {
	sc.rwMutex.RLock()
	defer sc.rwMutex.RUnlock()

	if sp, ok := sc.providers[key]; ok {
		return sp
	}
	return nil
}

func (sc *ServiceContainer) newInstance(sp ServiceProvider, params []any) (any, error) {
	if err := sp.Boot(sc); err != nil {
		return nil, err
	}
	if params == nil {
		params = sp.Params(sc)
	}

	inst, err := sp.Register(sc)(params...)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return inst, nil
}
