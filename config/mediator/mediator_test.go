package mediator

import (
	"reflect"
	"testing"
)

const (

	// PersistUrlRecord UrlRecord
	event1 Service = "event1"
	event2 Service = "event2"
	both   Service = "bouth"
)

type spy struct {
	arg    interface{}
	called bool
}

func (s *spy) function(v interface{}) error {
	s.called = true
	s.arg = v
	return nil
}

type myStruct struct {
	value string
}

func TestMediator(t *testing.T) {
	// Spy para o primeiro evento
	spyOne := &spy{}
	Get().Register(event1, spyOne.function)

	// Spy para o segundo evento
	spyTwo := &spy{}
	Get().Register(event2, spyTwo.function)

	// Spy para ambos os eventos
	spyBoth := &spy{}
	Get().Register(both, spyBoth.function)

	t.Run("call func one", func(t *testing.T) {
		Get().Notify(event1, myStruct{value: "test one"})
		if !spyOne.called {
			t.Error("expected handler for event1 to be called")
		}
		m := myStruct{value: "test one"}
		if spyOne.arg != m {
			t.Errorf("expected argument to be myStruct{value: \"test one\"}, got %+v", spyOne.arg)
		}
	})

	t.Run("call func two", func(t *testing.T) {
		data := &myStruct{value: "test two"}
		Get().Notify(event2, data)
		if !spyTwo.called {
			t.Error("expected handler for event2 to be called")
		}
		if !reflect.DeepEqual(spyTwo.arg, data) {
			t.Errorf("expected argument to be %+v, got %+v", data, spyTwo.arg)
		}
	})

	t.Run("call both", func(t *testing.T) {
		data := myStruct{value: "test both"}
		Get().Notify(both, data)
		if !spyBoth.called {
			t.Error("expected handler for both to be called")
		}
		if !reflect.DeepEqual(spyBoth.arg, data) {
			t.Errorf("expected argument to be %+v, got %+v", data, spyBoth.arg)
		}
	})
}
