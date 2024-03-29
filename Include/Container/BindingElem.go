package Container

import (
	"errors"
	"reflect"
)

type BindingImpl struct {
	concrete reflect.Value
	shared   bool
	alias    string
}

func NewBindingImpl(concrete interface{}) *BindingImpl {
	refValue := reflect.ValueOf(concrete)

	// check type.
	refType := refValue.Type().Kind()
	if refType != reflect.Ptr {
		panic(errors.New("UnSupport non-pointer implement type: " + refType.String()))
	}

	// combine struct.
	return &BindingImpl{
		concrete: refValue,
		shared:   false,
		alias:    "",
	}
}

func (e *BindingImpl) GetConcrete() reflect.Value {
	return e.concrete
}

func (e *BindingImpl) SetShared() *BindingImpl {
	e.shared = true
	return e
}

func (e *BindingImpl) GetShared() bool {
	return e.shared
}

func (e *BindingImpl) SetAlias(name string) *BindingImpl {
	e.alias = name
	return e
}

func (e *BindingImpl) GetAlias() string {
	return e.alias
}
