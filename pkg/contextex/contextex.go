package contextex

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrContextValuesNotFound    = errors.New("context values not found")
	ErrContextValuesKeyNotFound = errors.New("key not found")
	ErrTypeConversion           = errors.New("invalid type conversion")
)

// ContextKey comprises keys used to access common information from a request context.
type ContextKey string

// ContextMap - Context Map type wrapped by Values.
type ContextMap map[ContextKey]interface{}

// ContextValuesKey - Context Values Key type.
type ContextValuesKey string

// Values - struct to store values on context.
type Values struct {
	m ContextMap
}

func WrapError(message string, err error) error {
	return fmt.Errorf("%s: %w", message, err)
}

func NewValues(m ContextMap) Values {
	return Values{
		m: m,
	}
}

// GetMap - GetMap value from Values struct.
func (v Values) GetMap() ContextMap {
	return v.m
}

// GetKey - GetKey value from Values struct.
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

// GetFieldsMap - get field list from Context and append fields list to it.
func GetFieldsMap(ctx context.Context, contextValuesKey ContextValuesKey,
	fields ...interface{}) (map[string]interface{}, error) {
	results := make(map[string]interface{})

	ctxFields, ctxErr := GetContextValues(ctx, contextValuesKey)
	if ctxErr != nil {
		return nil, WrapError("GetFieldsMap->GetContextValues", ctxErr)
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

// GetContextValues contextValuesKey map stored in context.
func GetContextValues(ctx context.Context, contextValuesKey ContextValuesKey) (*Values, error) {
	values := ctx.Value(contextValuesKey)
	if values == nil {
		return nil, ErrContextValuesNotFound
	}

	if v, ok := values.(Values); ok {
		return &v, nil
	}

	return nil, ErrTypeConversion
}

// GetContextValuesKey get value for key from contextValuesKey map stored in context.
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

	end := 0
	if fields != nil {
		end = len(fields)
	}

	for idx := 0; idx < end; {
		var key string

		if k, ok := fields[idx].(string); ok {
			key = k
		} else {
			idx += 2

			continue
		}

		var value interface{}

		if (idx + 1) < end {
			value = fields[idx+1]
		}

		if valueType, ok := value.([]interface{}); ok {
			value = PairFields(valueType...)
		}

		results[ContextKey(key)] = value

		idx += 2
	}

	return results
}
