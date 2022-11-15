// Code generated .* DO NOT EDIT.
package gobdd

import "fmt"

func (ctx Context) GetString(key interface{}, defaultValue ...string) (string, error) {
	if len(defaultValue) > 1 {
		return "", fmt.Errorf("allowed to pass only 1 default value but %d got", len(defaultValue))
	}

	if _, ok := ctx.values[key]; !ok {
		if len(defaultValue) == 1 {
			return defaultValue[0], nil
		}
		return "", fmt.Errorf("the key %+v does not exist", key)
	}

	value, ok := ctx.values[key].(string)
	if !ok {
		return "", fmt.Errorf("the expected value is not string (%T)", key)
	}
	return value, nil
}

func (ctx Context) GetInt(key interface{}, defaultValue ...int) (int, error) {
	if len(defaultValue) > 1 {
		return 0, fmt.Errorf("allowed to pass only 1 default value but %d got", len(defaultValue))
	}

	if _, ok := ctx.values[key]; !ok {
		if len(defaultValue) == 1 {
			return defaultValue[0], nil
		}
		return 0, fmt.Errorf("the key %+v does not exist", key)
	}

	value, ok := ctx.values[key].(int)
	if !ok {
		return 0, fmt.Errorf("the expected value is not int (%T)", key)
	}
	return value, nil
}

func (ctx Context) GetInt8(key interface{}, defaultValue ...int8) (int8, error) {
	if len(defaultValue) > 1 {
		return 0, fmt.Errorf("allowed to pass only 1 default value but %d got", len(defaultValue))
	}

	if _, ok := ctx.values[key]; !ok {
		if len(defaultValue) == 1 {
			return defaultValue[0], nil
		}
		return 0, fmt.Errorf("the key %+v does not exist", key)
	}

	value, ok := ctx.values[key].(int8)
	if !ok {
		return 0, fmt.Errorf("the expected value is not int8 (%T)", key)
	}
	return value, nil
}

func (ctx Context) GetInt16(key interface{}, defaultValue ...int16) (int16, error) {
	if len(defaultValue) > 1 {
		return 0, fmt.Errorf("allowed to pass only 1 default value but %d got", len(defaultValue))
	}

	if _, ok := ctx.values[key]; !ok {
		if len(defaultValue) == 1 {
			return defaultValue[0], nil
		}
		return 0, fmt.Errorf("the key %+v does not exist", key)
	}

	value, ok := ctx.values[key].(int16)
	if !ok {
		return 0, fmt.Errorf("the expected value is not int16 (%T)", key)
	}
	return value, nil
}

func (ctx Context) GetInt32(key interface{}, defaultValue ...int32) (int32, error) {
	if len(defaultValue) > 1 {
		return 0, fmt.Errorf("allowed to pass only 1 default value but %d got", len(defaultValue))
	}

	if _, ok := ctx.values[key]; !ok {
		if len(defaultValue) == 1 {
			return defaultValue[0], nil
		}
		return 0, fmt.Errorf("the key %+v does not exist", key)
	}

	value, ok := ctx.values[key].(int32)
	if !ok {
		return 0, fmt.Errorf("the expected value is not int32 (%T)", key)
	}
	return value, nil
}

func (ctx Context) GetInt64(key interface{}, defaultValue ...int64) (int64, error) {
	if len(defaultValue) > 1 {
		return 0, fmt.Errorf("allowed to pass only 1 default value but %d got", len(defaultValue))
	}

	if _, ok := ctx.values[key]; !ok {
		if len(defaultValue) == 1 {
			return defaultValue[0], nil
		}
		return 0, fmt.Errorf("the key %+v does not exist", key)
	}

	value, ok := ctx.values[key].(int64)
	if !ok {
		return 0, fmt.Errorf("the expected value is not int64 (%T)", key)
	}
	return value, nil
}

func (ctx Context) GetFloat32(key interface{}, defaultValue ...float32) (float32, error) {
	if len(defaultValue) > 1 {
		return 0, fmt.Errorf("allowed to pass only 1 default value but %d got", len(defaultValue))
	}

	if _, ok := ctx.values[key]; !ok {
		if len(defaultValue) == 1 {
			return defaultValue[0], nil
		}
		return 0, fmt.Errorf("the key %+v does not exist", key)
	}

	value, ok := ctx.values[key].(float32)
	if !ok {
		return 0, fmt.Errorf("the expected value is not float32 (%T)", key)
	}
	return value, nil
}

func (ctx Context) GetFloat64(key interface{}, defaultValue ...float64) (float64, error) {
	if len(defaultValue) > 1 {
		return 0, fmt.Errorf("allowed to pass only 1 default value but %d got", len(defaultValue))
	}

	if _, ok := ctx.values[key]; !ok {
		if len(defaultValue) == 1 {
			return defaultValue[0], nil
		}
		return 0, fmt.Errorf("the key %+v does not exist", key)
	}

	value, ok := ctx.values[key].(float64)
	if !ok {
		return 0, fmt.Errorf("the expected value is not float64 (%T)", key)
	}
	return value, nil
}

func (ctx Context) GetBool(key interface{}, defaultValue ...bool) (bool, error) {
	if len(defaultValue) > 1 {
		return false, fmt.Errorf("allowed to pass only 1 default value but %d got", len(defaultValue))
	}

	if _, ok := ctx.values[key]; !ok {
		if len(defaultValue) == 1 {
			return defaultValue[0], nil
		}
		return false, fmt.Errorf("the key %+v does not exist", key)
	}

	value, ok := ctx.values[key].(bool)
	if !ok {
		return false, fmt.Errorf("the expected value is not bool (%T)", key)
	}
	return value, nil
}
