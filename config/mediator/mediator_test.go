package mediator

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

type event1 struct {
	arg    string
	called bool
}

func (e *event1) Notify(ctx context.Context, t any) error {
	e.arg = t.(myStruct).value
	e.called = true
	return nil
}

type event2 struct {
	arg    string
	called bool
}

func (e *event2) Notify(ctx context.Context, t any) error {
	e.arg = t.(myStruct).value
	e.called = true
	return nil
}

type both struct {
	arg    string
	called bool
}

func (e *both) Notify(ctx context.Context, t any) error {
	e.arg = t.(myStruct).value
	e.called = true
	return nil
}

type myStruct struct {
	value string
}

func TestMediator(t *testing.T) {
	spyOne := &event1{}
	spyTwo := &event2{}
	spyBoth := &both{}

	ctx := context.Background()
	t.Run("call func one", func(t *testing.T) {
		err := spyOne.Notify(ctx, myStruct{value: "test one"})
		assert.NoError(t, err)
		if !spyOne.called {
			t.Error("expected handler for event1 to be called")
		}
		m := myStruct{value: "test one"}
		if spyOne.arg != m.value {
			t.Errorf("expected argument to be myStruct{value: \"test one\"}, got %+v", spyOne.arg)
		}
	})

	t.Run("call func two", func(t *testing.T) {
		data := myStruct{value: "test two"}
		err := spyTwo.Notify(ctx, data)
		assert.NoError(t, err)
		if !spyTwo.called {
			t.Error("expected handler for event2 to be called")
		}
		if spyTwo.arg != data.value {
			t.Errorf("expected argument to be myStruct{value: \"test one\"}, got %+v", spyOne.arg)
		}
	})

	t.Run("call both", func(t *testing.T) {
		data := myStruct{value: "test two"}
		err := spyBoth.Notify(ctx, data)
		assert.NoError(t, err)
		if !spyBoth.called {
			t.Error("expected handler for both to be called")
		}
		if spyBoth.arg != data.value {
			t.Errorf("expected argument to be myStruct{value: \"test one\"}, got %+v", spyOne.arg)
		}
	})
}
