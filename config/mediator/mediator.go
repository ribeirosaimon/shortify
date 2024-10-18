package mediator

import (
	"fmt"

	"github.com/ribeirosaimon/tooltip/tlog"
)

var mediator *Mediator

type Mediator struct {
	dependencies map[Service][]*dependency
}

type dependency struct {
	function func(interface{}) error
}

func Get() *Mediator {
	if mediator == nil {
		mediator = &Mediator{
			dependencies: make(map[Service][]*dependency),
		}
	}
	return mediator
}

func (m *Mediator) Register(event Service, callback func(T any) error) {
	m.dependencies[event] = append(m.dependencies[event], &dependency{
		function: callback,
	})
}

func (m *Mediator) Notify(event Service, data interface{}) {
	if dependencies, exists := m.dependencies[event]; exists {
		for _, dep := range dependencies {
			// Remover a chamada dep.once.Do e executar diretamente
			if err := dep.function(data); err != nil {
				tlog.Error("Mediator.Notify", fmt.Sprintf("Handle error in dependency %p with data %s", dep.function, data))
			}
		}
	}
}
