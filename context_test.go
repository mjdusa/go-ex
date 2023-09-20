package ext

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewValues(t *testing.T) {
	empty := make(ContextMap)
	vals := make(ContextMap)
	vals["foo"] = "bar"

	type provided struct {
		Data *ContextMap
	}
	type expected struct {
		Error error
		Value *ContextMap
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
		actual := NewValues(*test.Provided.Data)

		assert.Equal(t, len(*test.Expected.Value), len(actual.m), test.Name)

		for key, val := range *test.Expected.Value {
			found := actual.m[key]
			assert.Equal(t, val, found, test.Name)
		}
	}
}

func Test_Values_GetMap(t *testing.T) {
	empty := make(ContextMap)
	vals := make(ContextMap)
	vals["foo"] = "bar"

	type provided struct {
		Data *ContextMap
	}
	type expected struct {
		Error error
		Value *ContextMap
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
		actual := NewValues(*test.Provided.Data)

		assert.Equal(t, len(*test.Expected.Value), len(actual.GetMap()), test.Name)

		for key, val := range *test.Expected.Value {
			found := actual.GetMap()[key]
			assert.Equal(t, val, found, test.Name)
		}
	}
}

func Test_Values_GetKey(t *testing.T) {
	empty := make(ContextMap)
	vals := make(ContextMap)
	vals["foo"] = "bar"

	type provided struct {
		Data *ContextMap
	}
	type expected struct {
		Error error
		Value *ContextMap
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
		actual := NewValues(*test.Provided.Data)

		assert.Equal(t, len(*test.Expected.Value), len(actual.GetMap()), test.Name)

		for key, val := range *test.Expected.Value {
			found := actual.GetKey(key)
			assert.Equal(t, val, found, test.Name)
		}
	}
}

func Test_CreateContextWithValues(t *testing.T) {
	empty := make(ContextMap)
	vals := make(ContextMap)
	vals["foo"] = "bar"

	type provided struct {
		ValuesKey ContextValuesKey
		Data      *ContextMap
	}
	type expected struct {
		Error error
		Value *ContextMap
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
		values := NewValues(*test.Provided.Data)
		ctx := context.Background()
		ctx = CreateContextWithValues(ctx, test.Provided.ValuesKey, values)

		actualValues, actualErr := GetContextValues(ctx, test.Provided.ValuesKey)

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

func Test_CreateContextWithMap(t *testing.T) {
	empty := make(ContextMap)
	vals := make(ContextMap)
	vals["foo"] = "bar"

	type provided struct {
		ValuesKey ContextValuesKey
		Data      *ContextMap
	}
	type expected struct {
		Error error
		Value *ContextMap
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
		ctx = CreateContextWithMap(ctx, test.Provided.ValuesKey, *test.Provided.Data)

		actualValues, actualErr := GetContextValues(ctx, test.Provided.ValuesKey)

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

func Test_GetFieldsMap(t *testing.T) {
	key0 := "key-0"
	value0 := "value-0"
	key1 := "key-1"
	value1 := "value-1"
	key2 := "key-2"
	value2 := []interface{}{"sub-key-2", "sub-value-2"}
	fields := []interface{}{key1, value1, key2, value2}

	emptyResponse := make(map[string]interface{})
	fieldsOnlyResponse := make(map[string]interface{})
	mapOnlyResponse := make(map[string]interface{})
	bothResponse := make(map[string]interface{})

	fieldsOnlyResponse[key1] = value1
	fieldsOnlyResponse[key2] = value2
	mapOnlyResponse[key0] = value0
	bothResponse[key0] = value0
	bothResponse[key1] = value1
	bothResponse[key2] = value2

	emptyMap := make(ContextMap)
	valsMap := make(ContextMap)
	valsMap[ContextKey(key0)] = value0

	type provided struct {
		ValuesKey  ContextValuesKey
		ContextMap *ContextMap
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
				ValuesKey:  ContextValuesKey("valuesKey"),
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
				ValuesKey:  ContextValuesKey("valuesKey"),
				ContextMap: &emptyMap,
				Fields:     fields,
			},
			Expected: expected{
				Error: nil,
				Value: fieldsOnlyResponse,
			},
		},
		{
			Name: "Test populated map, nil fields",
			Provided: provided{
				ValuesKey:  ContextValuesKey("valuesKey"),
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
				ValuesKey:  ContextValuesKey("valuesKey"),
				ContextMap: &valsMap,
				Fields:     fields,
			},
			Expected: expected{
				Error: nil,
				Value: bothResponse,
			},
		},
	}

	for _, test := range tests {
		ctx := context.Background()
		ctx = CreateContextWithMap(ctx, test.Provided.ValuesKey, *test.Provided.ContextMap)

		actualValue, actualErr := GetFieldsMap(ctx, test.Provided.ValuesKey, test.Provided.Fields...)

		assert.Equal(t, test.Expected.Error, actualErr, test.Name)
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

func Test_GetFieldsMap_Error(t *testing.T) {
	emptyMap := make(ContextMap)

	type provided struct {
		ValuesKey  ContextValuesKey
		ContextMap *ContextMap
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
				ValuesKey:  ContextValuesKey("valuesKey"),
				ContextMap: &emptyMap,
				Fields:     nil,
			},
			Expected: expected{
				Error: ErrContextValuesNotFound,
				Value: nil,
			},
		},
	}

	for _, test := range tests {
		ctx := context.Background()

		actualValue, actualErr := GetFieldsMap(ctx, test.Provided.ValuesKey, test.Provided.Fields...)

		assert.Equal(t, test.Expected.Error, actualErr, test.Name)
		assert.Nil(t, actualValue, test.Name)
	}
}

func Test_GetContextValues(t *testing.T) {
}

func Test_GetContextValuesKey(t *testing.T) {
	key := ContextKey("key")
	value := "value"
	m := make(ContextMap)
	m[key] = value

	type provided struct {
		ValuesKey  ContextValuesKey
		ContextMap *ContextMap
		Key        ContextKey
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
				ValuesKey:  ContextValuesKey("Values"),
				ContextMap: &m,
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
		ctx = CreateContextWithMap(ctx, test.Provided.ValuesKey, *test.Provided.ContextMap)

		actualValue, actualErr := GetContextValuesKey(ctx, test.Provided.ValuesKey, test.Provided.Key)

		assert.Equal(t, test.Expected.Error, actualErr, test.Name)
		assert.Equal(t, test.Expected.Value, actualValue, test.Name)
	}
}

func Test_GetContextValuesKey_Error(t *testing.T) {
	emptyMap := make(ContextMap)

	type provided struct {
		ValuesKey  ContextValuesKey
		ContextMap *ContextMap
		Key        ContextKey
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
				ValuesKey:  ContextValuesKey(""),
				ContextMap: &emptyMap,
				Key:        ContextKey(""),
			},
			Expected: expected{
				Error: ErrContextValuesNotFound,
				Value: nil,
			},
		},
		{
			Name: "Test missing Key error",
			Provided: provided{
				ValuesKey:  ContextValuesKey("valuesKey"),
				ContextMap: &emptyMap,
				Key:        ContextKey("key"),
			},
			Expected: expected{
				Error: ErrContextValuesKeyNotFound,
				Value: nil,
			},
		},
	}

	for _, test := range tests {
		ctx := context.Background()
		if len(test.Provided.ValuesKey) > 0 {
			cm := make(ContextMap)
			ctx = CreateContextWithMap(ctx, test.Provided.ValuesKey, cm)
		}

		actualValue, actualErr := GetContextValuesKey(ctx, test.Provided.ValuesKey, test.Provided.Key)

		assert.Equal(t, test.Expected.Error, actualErr, test.Name)
		assert.Nil(t, actualValue, test.Name)
	}
}
