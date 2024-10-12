package di

import "sync"

var registry *Registry

type Registry struct {
	dependencies map[Service]*dependency
}

type dependency struct {
	instance any
	once     sync.Once
}

func GetRegistry() *Registry {
	if registry == nil {
		registry = &Registry{
			dependencies: make(map[Service]*dependency),
		}
	}
	return registry
}

func (r *Registry) Provide(name Service, factoryFunc func() any) {
	if _, exists := r.dependencies[name]; !exists {
		r.dependencies[name] = &dependency{}
	}

	r.dependencies[name].once.Do(func() {
		r.dependencies[name].instance = factoryFunc()
	})

}

func (r *Registry) Inject(name Service) any {
	return r.dependencies[name].instance
}
