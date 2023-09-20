package ext

import (
	"context"
	"errors"
)

var (
	ErrContextValuesNotFound    = errors.New("context values not found")
	ErrContextValuesKeyNotFound = errors.New("key not found")
)

// ContextKey comprises keys used to access common information from a request context.
type ContextKey string

type ContextMap map[ContextKey]interface{}

// ContextValuesKey - ...
type ContextValuesKey string

// Values - struct to store values on context
type Values struct {
	m ContextMap
}

func NewValues(m ContextMap) Values {
	return Values{
		m: m,
	}
}

// Get - Get value from Values struct
func (v Values) GetMap() ContextMap {
	return v.m
}

// Get - Get value from Values struct
func (v Values) GetKey(key ContextKey) interface{} {
	return v.m[key]
}

func CreateContextWithMap(ctx context.Context, valuesKey ContextValuesKey, cm ContextMap) context.Context {
	values := NewValues(cm)
	return CreateContextWithValues(ctx, valuesKey, values)
}

func CreateContextWithValues(ctx context.Context, valuesKey ContextValuesKey, values Values) context.Context {
	return context.WithValue(ctx, valuesKey, values)
}

// GetFieldsMap - get
func GetFieldsMap(ctx context.Context, contextValuesKey ContextValuesKey, fields ...interface{}) (map[string]interface{}, error) {
	results := make(map[string]interface{})

	ctxFields, ctxErr := GetContextValues(ctx, contextValuesKey)
	if ctxErr != nil {
		return nil, ctxErr
	}

	for k, v := range ctxFields.GetMap() {
		results[string(k)] = v
	}

	if len(fields) > 0 {
		fm := PairFields(fields...)
		for k, v := range fm {
			results[string(k)] = v
		}
	}

	return results, nil
}

// GetContextValues contextValuesKey map stored in context
func GetContextValues(ctx context.Context, contextValuesKey ContextValuesKey) (*Values, error) {
	values := ctx.Value(contextValuesKey)
	if values == nil {
		return nil, ErrContextValuesNotFound
	}

	v := values.(Values)

	return &v, nil
}

// GetContextValuesKey get value for key from contextValuesKey map stored in context
func GetContextValuesKey(ctx context.Context, contextValuesKey ContextValuesKey, key ContextKey) (interface{}, error) {
	values, err := GetContextValues(ctx, contextValuesKey)
	if err != nil {
		return nil, err
	}

	value := values.GetKey(key)
	if value == nil {
		return nil, ErrContextValuesKeyNotFound
	}

	return value, nil
}

func PairFields(fields ...interface{}) ContextMap {
	results := make(ContextMap)

	end := len(fields)

	for idx := 0; idx < end; {
		key := fields[idx].(string)

		var value interface{}

		if (idx + 1) < end {
			value = fields[idx+1]
		}

		switch valueType := value.(type) {
		case []interface{}:
			value = PairFields(valueType...)
		}

		results[ContextKey(key)] = value

		idx += 2
	}

	return results
}
