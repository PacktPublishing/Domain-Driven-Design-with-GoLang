package gobdd

import (
	"encoding/json"
	"fmt"
)

// Holds data from previously executed steps
type Context struct {
	values map[interface{}]interface{}
}

// Creates a new (empty) context struct
func NewContext() Context {
	return Context{
		values: map[interface{}]interface{}{},
	}
}

// Clone creates a copy of the context
func (ctx Context) Clone() Context {
	c := Context{
		values: map[interface{}]interface{}{},
	}

	for k, v := range ctx.values {
		c.Set(k, v)
	}

	return c
}

// Sets the value under the key
func (ctx Context) Set(key interface{}, value interface{}) {
	ctx.values[key] = value
}

// Returns the data under the key.
// If couldn't find anything but the default value is provided, returns the default value.
// Otherwise, it returns an error.
func (ctx Context) Get(key interface{}, defaultValue ...interface{}) (interface{}, error) {
	if _, ok := ctx.values[key]; !ok {
		if len(defaultValue) == 1 {
			return defaultValue[0], nil
		}

		return nil, fmt.Errorf("the key %+v does not exist", key)
	}

	return ctx.values[key], nil
}

// GetAs copies data from tke key to the dest.
// Supports maps, slices and structs.
func (ctx Context) GetAs(key interface{}, dest interface{}) error {
	if _, ok := ctx.values[key]; !ok {
		return fmt.Errorf("the key %+v does not exist", key)
	}

	d, err := json.Marshal(ctx.values[key])
	if err != nil {
		return err
	}

	return json.Unmarshal(d, dest)
}

// It is a shortcut for getting the value already casted as error.
func (ctx Context) GetError(key interface{}, defaultValue ...error) (error, error) {
	if _, ok := ctx.values[key]; !ok {
		if len(defaultValue) == 1 {
			return defaultValue[0], nil
		}

		return nil, fmt.Errorf("the key %+v does not exist", key)
	}

	if ctx.values[key] == nil {
		return nil, nil
	}

	value, ok := ctx.values[key].(error)
	if !ok {
		return nil, fmt.Errorf("the expected value is not error  (%T)", key)
	}

	return value, nil
}
