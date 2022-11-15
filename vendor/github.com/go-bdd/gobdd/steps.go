package gobdd

import (
	"errors"
	"reflect"
)

func validateStepFunc(f interface{}) error {
	value := reflect.ValueOf(f)
	if value.Kind() != reflect.Func {
		return errors.New("the parameter should be a function")
	}

	if value.Type().NumIn() < 2 {
		return errors.New("the function should have StepTest and Context as the first argument")
	}

	val := value.Type().In(0)

	testingInterface := reflect.TypeOf((*StepTest)(nil)).Elem()
	if !val.Implements(testingInterface) {
		return errors.New("the function should have the StepTest as the first argument")
	}

	val = value.Type().In(1)

	n := val.ConvertibleTo(reflect.TypeOf((*Context)(nil)).Elem())
	if !n {
		return errors.New("the function should have Context as the second argument")
	}

	return nil
}
