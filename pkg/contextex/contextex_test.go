package contextex_test

import (
	"context"
	"testing"

	"github.com/mjdusa/go-ext/pkg/contextex"
	"github.com/stretchr/testify/assert"
)

/*
func TestNewValues(t *testing.T) {
	empty := make(contextex.ContextMap)
	vals := make(contextex.ContextMap)
	vals["foo"] = "bar"

	type provided struct {
		Data *contextex.ContextMap
	}
	type expected struct {
		Error error
		Value *contextex.ContextMap
	}

	tests := []struct {
		Name     string
		Provided provided
		Expected expected
	}{
		{
			Name: "Test empty map",
			Provided: provided{
				Data: &empty,
			},
			Expected: expected{
				Error: nil,
				Value: &empty,
			},
		},
		{
			Name: "Test populated map",
			Provided: provided{
				Data: &vals,
			},
			Expected: expected{
				Error: nil,
				Value: &vals,
			},
		},
	}

	for _, test := range tests {
		actual := contextex.NewValues(*test.Provided.Data)

		assert.Equal(t, len(*test.Expected.Value), len(actual.m), test.Name)

		for key, val := range *test.Expected.Value {
			found := actual.m[key]
			assert.Equal(t, val, found, test.Name)
		}
	}
}
*/

func TestGetMap(t *testing.T) {
	empty := make(contextex.ContextMap)
	vals := make(contextex.ContextMap)
	vals["foo"] = "bar"

	type provided struct {
		Data *contextex.ContextMap
	}
	type expected struct {
		Error error
		Value *contextex.ContextMap
	}

	tests := []struct {
		Name     string
		Provided provided
		Expected expected
	}{
		{
			Name: "Test empty map",
			Provided: provided{
				Data: &empty,
			},
			Expected: expected{
				Error: nil,
				Value: &empty,
			},
		},
		{
			Name: "Test populated map",
			Provided: provided{
				Data: &vals,
			},
			Expected: expected{
				Error: nil,
				Value: &vals,
			},
		},
	}

	for _, test := range tests {
		actual := contextex.NewValues(*test.Provided.Data)

		assert.Equal(t, len(*test.Expected.Value), len(actual.GetMap()), test.Name)

		for key, val := range *test.Expected.Value {
			found := actual.GetMap()[key]
			assert.Equal(t, val, found, test.Name)
		}
	}
}

func TestGetKey(t *testing.T) {
	empty := make(contextex.ContextMap)
	vals := make(contextex.ContextMap)
	vals["foo"] = "bar"

	type provided struct {
		Data *contextex.ContextMap
	}
	type expected struct {
		Error error
		Value *contextex.ContextMap
	}

	tests := []struct {
		Name     string
		Provided provided
		Expected expected
	}{
		{
			Name: "Test empty map",
			Provided: provided{
				Data: &empty,
			},
			Expected: expected{
				Error: nil,
				Value: &empty,
			},
		},
		{
			Name: "Test populated map",
			Provided: provided{
				Data: &vals,
			},
			Expected: expected{
				Error: nil,
				Value: &vals,
			},
		},
	}

	for _, test := range tests {
		actual := contextex.NewValues(*test.Provided.Data)

		assert.Equal(t, len(*test.Expected.Value), len(actual.GetMap()), test.Name)

		for key, val := range *test.Expected.Value {
			found := actual.GetKey(key)
			assert.Equal(t, val, found, test.Name)
		}
	}
}

func TestCreateContextWithValues(t *testing.T) {
	empty := make(contextex.ContextMap)
	vals := make(contextex.ContextMap)
	vals["foo"] = "bar"

	type provided struct {
		ValuesKey contextex.ContextValuesKey
		Data      *contextex.ContextMap
	}
	type expected struct {
		Error error
		Value *contextex.ContextMap
	}

	tests := []struct {
		Name     string
		Provided provided
		Expected expected
	}{
		{
			Name: "Test empty map",
			Provided: provided{
				ValuesKey: "valuesKey",
				Data:      &empty,
			},
			Expected: expected{
				Error: nil,
				Value: &empty,
			},
		},
		{
			Name: "Test populated map",
			Provided: provided{
				ValuesKey: "valuesKey",
				Data:      &vals,
			},
			Expected: expected{
				Error: nil,
				Value: &vals,
			},
		},
	}

	for _, test := range tests {
		values := contextex.NewValues(*test.Provided.Data)
		ctx := context.Background()
		ctx = contextex.CreateContextWithValues(ctx, test.Provided.ValuesKey, values)

		actualValues, actualErr := contextex.GetContextValues(ctx, test.Provided.ValuesKey)

		assert.Equal(t, test.Expected.Error, actualErr, test.Name)
		if actualValues != nil {
			assert.NotNil(t, test.Expected.Value, test.Name)
			assert.Equal(t, len(*test.Expected.Value), len(actualValues.GetMap()), test.Name)

			for key, val := range *test.Expected.Value {
				found := actualValues.GetKey(key)
				assert.Equal(t, val, found, test.Name)
			}
		} else {
			assert.Nil(t, test.Expected.Value, test.Name)
		}
	}
}

func TestCreateContextWithMap(t *testing.T) {
	empty := make(contextex.ContextMap)
	vals := make(contextex.ContextMap)
	vals["foo"] = "bar"

	type provided struct {
		ValuesKey contextex.ContextValuesKey
		Data      *contextex.ContextMap
	}
	type expected struct {
		Error error
		Value *contextex.ContextMap
	}

	tests := []struct {
		Name     string
		Provided provided
		Expected expected
	}{
		{
			Name: "Test empty map",
			Provided: provided{
				ValuesKey: "valuesKey",
				Data:      &empty,
			},
			Expected: expected{
				Error: nil,
				Value: &empty,
			},
		},
		{
			Name: "Test populated map",
			Provided: provided{
				ValuesKey: "valuesKey",
				Data:      &vals,
			},
			Expected: expected{
				Error: nil,
				Value: &vals,
			},
		},
	}

	for _, test := range tests {
		ctx := context.Background()
		ctx = contextex.CreateContextWithMap(ctx, test.Provided.ValuesKey, *test.Provided.Data)

		actualValues, actualErr := contextex.GetContextValues(ctx, test.Provided.ValuesKey)

		assert.Equal(t, test.Expected.Error, actualErr, test.Name)
		if actualValues != nil {
			assert.NotNil(t, test.Expected.Value, test.Name)
			assert.Equal(t, len(*test.Expected.Value), len(actualValues.GetMap()), test.Name)

			for key, val := range *test.Expected.Value {
				found := actualValues.GetKey(key)
				assert.Equal(t, val, found, test.Name)
			}
		} else {
			assert.Nil(t, test.Expected.Value, test.Name)
		}
	}
}

func TestGetFieldsMap(t *testing.T) {
	key0 := "key-0"
	value0 := "value-0"
	key1 := "key-1"
	value1 := "value-1"
	key2 := "key-2"
	value2 := []interface{}{"sub-key-2", "sub-value-2"}
	goodFields := []interface{}{key1, value1, key2, value2}
	badKeyFeilds := []interface{}{999, value1, key2, value2}

	emptyResponse := make(map[string]interface{})
	fieldsOnlyResponse := make(map[string]interface{})
	mapOnlyResponse := make(map[string]interface{})
	bothResponse := make(map[string]interface{})
	badKeyResponse := make(map[string]interface{})

	fieldsOnlyResponse[key1] = value1
	fieldsOnlyResponse[key2] = value2
	mapOnlyResponse[key0] = value0
	bothResponse[key0] = value0
	bothResponse[key1] = value1
	bothResponse[key2] = value2
	badKeyResponse[key0] = value0
	badKeyResponse[key2] = value2

	emptyMap := make(contextex.ContextMap)
	valsMap := make(contextex.ContextMap)
	valsMap[contextex.ContextKey(key0)] = value0

	type provided struct {
		ValuesKey  contextex.ContextValuesKey
		ContextMap *contextex.ContextMap
		Fields     []interface{}
	}
	type expected struct {
		Error error
		Value map[string]interface{}
	}

	tests := []struct {
		Name     string
		Provided provided
		Expected expected
	}{
		{
			Name: "Test empty map, nil fields",
			Provided: provided{
				ValuesKey:  contextex.ContextValuesKey("valuesKey"),
				ContextMap: &emptyMap,
				Fields:     nil,
			},
			Expected: expected{
				Error: nil,
				Value: emptyResponse,
			},
		},
		{
			Name: "Test empty map, populated fields",
			Provided: provided{
				ValuesKey:  contextex.ContextValuesKey("valuesKey"),
				ContextMap: &emptyMap,
				Fields:     goodFields,
			},
			Expected: expected{
				Error: nil,
				Value: fieldsOnlyResponse,
			},
		},
		{
			Name: "Test populated map, nil fields",
			Provided: provided{
				ValuesKey:  contextex.ContextValuesKey("valuesKey"),
				ContextMap: &valsMap,
				Fields:     nil,
			},
			Expected: expected{
				Error: nil,
				Value: mapOnlyResponse,
			},
		},
		{
			Name: "Test populated map, populated fields",
			Provided: provided{
				ValuesKey:  contextex.ContextValuesKey("valuesKey"),
				ContextMap: &valsMap,
				Fields:     goodFields,
			},
			Expected: expected{
				Error: nil,
				Value: bothResponse,
			},
		},
		{
			Name: "Test populated map, populated fields",
			Provided: provided{
				ValuesKey:  contextex.ContextValuesKey("valuesKey"),
				ContextMap: &valsMap,
				Fields:     badKeyFeilds,
			},
			Expected: expected{
				Error: nil,
				Value: badKeyResponse,
			},
		},
	}

	for _, test := range tests {
		ctx := context.Background()
		ctx = contextex.CreateContextWithMap(ctx, test.Provided.ValuesKey, *test.Provided.ContextMap)

		actualValue, actualErr := contextex.GetFieldsMap(ctx, test.Provided.ValuesKey, test.Provided.Fields...)

		assert.Equal(t, test.Expected.Error, actualErr, test.Name)
		//nolint:nestif  // unit test
		if actualValue == nil {
			assert.Nil(t, test.Expected.Value, test.Name)
		} else {
			assert.NotNil(t, test.Expected.Value, test.Name)
			assert.Equal(t, len(test.Expected.Value), len(actualValue), test.Name)

			for key, val := range test.Expected.Value {
				found := actualValue[key]
				if f, ok := found.(string); ok {
					if v, ok := val.(string); ok {
						assert.Equal(t, v, f, test.Name)
					} else {
						assert.Fail(t, "val type mismatch", test.Name)
					}
				}
			}
		}
	}
}

func TestGetFieldsMapError(t *testing.T) {
	emptyMap := make(contextex.ContextMap)

	type provided struct {
		ValuesKey  contextex.ContextValuesKey
		ContextMap *contextex.ContextMap
		Fields     []interface{}
	}
	type expected struct {
		Error error
		Value map[string]interface{}
	}

	tests := []struct {
		Name     string
		Provided provided
		Expected expected
	}{
		{
			Name: "Test missing Values error",
			Provided: provided{
				ValuesKey:  contextex.ContextValuesKey("valuesKey"),
				ContextMap: &emptyMap,
				Fields:     nil,
			},
			Expected: expected{
				Error: contextex.WrapError("GetFieldsMap->GetContextValues", contextex.ErrContextValuesNotFound),
				Value: nil,
			},
		},
	}

	for _, test := range tests {
		ctx := context.Background()

		actualValue, actualErr := contextex.GetFieldsMap(ctx, test.Provided.ValuesKey, test.Provided.Fields...)

		assert.Equal(t, test.Expected.Error.Error(), actualErr.Error(), test.Name)
		assert.Nil(t, actualValue, test.Name)
	}
}

func TestGetContextValues(t *testing.T) {
	emptyMap := make(contextex.ContextMap)

	type provided struct {
		ValuesKey  contextex.ContextValuesKey
		ContextMap *contextex.ContextMap
		Fields     []interface{}
	}
	type expected struct {
		Error error
		Value map[string]interface{}
	}

	tests := []struct {
		Name     string
		Provided provided
		Expected expected
	}{
		{
			Name: "Test missing Values error",
			Provided: provided{
				ValuesKey:  contextex.ContextValuesKey("valuesKey"),
				ContextMap: &emptyMap,
				Fields:     nil,
			},
			Expected: expected{
				Error: contextex.ErrContextValuesNotFound,
				Value: nil,
			},
		},
	}

	for _, test := range tests {
		ctx := context.Background()

		actualValue, actualErr := contextex.GetContextValues(ctx, test.Provided.ValuesKey)

		assert.Equal(t, test.Expected.Error.Error(), actualErr.Error(), test.Name)
		assert.Nil(t, actualValue, test.Name)
	}
}

func TestGetContextValuesKey(t *testing.T) {
	key := contextex.ContextKey("key")
	value := "value"
	cMap := make(contextex.ContextMap)
	cMap[key] = value

	type provided struct {
		ValuesKey  contextex.ContextValuesKey
		ContextMap *contextex.ContextMap
		Key        contextex.ContextKey
	}
	type expected struct {
		Error error
		Value interface{}
	}

	tests := []struct {
		Name     string
		Provided provided
		Expected expected
	}{
		{
			Name: "Test missing Values error",
			Provided: provided{
				ValuesKey:  contextex.ContextValuesKey("Values"),
				ContextMap: &cMap,
				Key:        key,
			},
			Expected: expected{
				Error: nil,
				Value: value,
			},
		},
	}

	for _, test := range tests {
		ctx := context.Background()
		ctx = contextex.CreateContextWithMap(ctx, test.Provided.ValuesKey, *test.Provided.ContextMap)

		actualValue, actualErr := contextex.GetContextValuesKey(ctx, test.Provided.ValuesKey, test.Provided.Key)

		assert.Equal(t, test.Expected.Error, actualErr, test.Name)
		assert.Equal(t, test.Expected.Value, actualValue, test.Name)
	}
}

func TestGetContextValuesKeyError(t *testing.T) {
	emptyMap := make(contextex.ContextMap)

	type provided struct {
		ValuesKey  contextex.ContextValuesKey
		ContextMap *contextex.ContextMap
		Key        contextex.ContextKey
	}
	type expected struct {
		Error error
		Value map[string]interface{}
	}

	tests := []struct {
		Name     string
		Provided provided
		Expected expected
	}{
		{
			Name: "Test missing Values error",
			Provided: provided{
				ValuesKey:  contextex.ContextValuesKey(""),
				ContextMap: &emptyMap,
				Key:        contextex.ContextKey(""),
			},
			Expected: expected{
				Error: contextex.ErrContextValuesNotFound,
				Value: nil,
			},
		},
		{
			Name: "Test missing Key error",
			Provided: provided{
				ValuesKey:  contextex.ContextValuesKey("valuesKey"),
				ContextMap: &emptyMap,
				Key:        contextex.ContextKey("key"),
			},
			Expected: expected{
				Error: contextex.ErrContextValuesKeyNotFound,
				Value: nil,
			},
		},
	}

	for _, test := range tests {
		ctx := context.Background()
		if len(test.Provided.ValuesKey) > 0 {
			cm := make(contextex.ContextMap)
			ctx = contextex.CreateContextWithMap(ctx, test.Provided.ValuesKey, cm)
		}

		actualValue, actualErr := contextex.GetContextValuesKey(ctx, test.Provided.ValuesKey, test.Provided.Key)

		assert.Equal(t, test.Expected.Error, actualErr, test.Name)
		assert.Nil(t, actualValue, test.Name)
	}
}
