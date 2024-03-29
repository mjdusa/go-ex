package ext_test

import (
	"context"
	"testing"

	"github.com/mjdusa/go-ext"
	"github.com/stretchr/testify/assert"
)

/*
func TestNewValues(t *testing.T) {
	empty := make(ext.ContextMap)
	vals := make(ext.ContextMap)
	vals["foo"] = "bar"

	type provided struct {
		Data *ext.ContextMap
	}
	type expected struct {
		Error error
		Value *ext.ContextMap
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
		actual := ext.NewValues(*test.Provided.Data)

		assert.Equal(t, len(*test.Expected.Value), len(actual.m), test.Name)

		for key, val := range *test.Expected.Value {
			found := actual.m[key]
			assert.Equal(t, val, found, test.Name)
		}
	}
}
*/

func TestGetMap(t *testing.T) {
	empty := make(ext.ContextMap)
	vals := make(ext.ContextMap)
	vals["foo"] = "bar"

	type provided struct {
		Data *ext.ContextMap
	}
	type expected struct {
		Error error
		Value *ext.ContextMap
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
		actual := ext.NewValues(*test.Provided.Data)

		assert.Equal(t, len(*test.Expected.Value), len(actual.GetMap()), test.Name)

		for key, val := range *test.Expected.Value {
			found := actual.GetMap()[key]
			assert.Equal(t, val, found, test.Name)
		}
	}
}

func TestGetKey(t *testing.T) {
	empty := make(ext.ContextMap)
	vals := make(ext.ContextMap)
	vals["foo"] = "bar"

	type provided struct {
		Data *ext.ContextMap
	}
	type expected struct {
		Error error
		Value *ext.ContextMap
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
		actual := ext.NewValues(*test.Provided.Data)

		assert.Equal(t, len(*test.Expected.Value), len(actual.GetMap()), test.Name)

		for key, val := range *test.Expected.Value {
			found := actual.GetKey(key)
			assert.Equal(t, val, found, test.Name)
		}
	}
}

func TestCreateContextWithValues(t *testing.T) {
	empty := make(ext.ContextMap)
	vals := make(ext.ContextMap)
	vals["foo"] = "bar"

	type provided struct {
		ValuesKey ext.ContextValuesKey
		Data      *ext.ContextMap
	}
	type expected struct {
		Error error
		Value *ext.ContextMap
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
		values := ext.NewValues(*test.Provided.Data)
		ctx := context.Background()
		ctx = ext.CreateContextWithValues(ctx, test.Provided.ValuesKey, values)

		actualValues, actualErr := ext.GetContextValues(ctx, test.Provided.ValuesKey)

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
	empty := make(ext.ContextMap)
	vals := make(ext.ContextMap)
	vals["foo"] = "bar"

	type provided struct {
		ValuesKey ext.ContextValuesKey
		Data      *ext.ContextMap
	}
	type expected struct {
		Error error
		Value *ext.ContextMap
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
		ctx = ext.CreateContextWithMap(ctx, test.Provided.ValuesKey, *test.Provided.Data)

		actualValues, actualErr := ext.GetContextValues(ctx, test.Provided.ValuesKey)

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

	emptyMap := make(ext.ContextMap)
	valsMap := make(ext.ContextMap)
	valsMap[ext.ContextKey(key0)] = value0

	type provided struct {
		ValuesKey  ext.ContextValuesKey
		ContextMap *ext.ContextMap
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
				ValuesKey:  ext.ContextValuesKey("valuesKey"),
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
				ValuesKey:  ext.ContextValuesKey("valuesKey"),
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
				ValuesKey:  ext.ContextValuesKey("valuesKey"),
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
				ValuesKey:  ext.ContextValuesKey("valuesKey"),
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
				ValuesKey:  ext.ContextValuesKey("valuesKey"),
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
		ctx = ext.CreateContextWithMap(ctx, test.Provided.ValuesKey, *test.Provided.ContextMap)

		actualValue, actualErr := ext.GetFieldsMap(ctx, test.Provided.ValuesKey, test.Provided.Fields...)

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
	emptyMap := make(ext.ContextMap)

	type provided struct {
		ValuesKey  ext.ContextValuesKey
		ContextMap *ext.ContextMap
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
				ValuesKey:  ext.ContextValuesKey("valuesKey"),
				ContextMap: &emptyMap,
				Fields:     nil,
			},
			Expected: expected{
				Error: ext.WrapError("GetFieldsMap->GetContextValues", ext.ErrContextValuesNotFound),
				Value: nil,
			},
		},
	}

	for _, test := range tests {
		ctx := context.Background()

		actualValue, actualErr := ext.GetFieldsMap(ctx, test.Provided.ValuesKey, test.Provided.Fields...)

		assert.Equal(t, test.Expected.Error.Error(), actualErr.Error(), test.Name)
		assert.Nil(t, actualValue, test.Name)
	}
}

func TestGetContextValues(t *testing.T) {
	emptyMap := make(ext.ContextMap)

	type provided struct {
		ValuesKey  ext.ContextValuesKey
		ContextMap *ext.ContextMap
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
				ValuesKey:  ext.ContextValuesKey("valuesKey"),
				ContextMap: &emptyMap,
				Fields:     nil,
			},
			Expected: expected{
				Error: ext.WrapError("", ext.ErrContextValuesNotFound),
				Value: nil,
			},
		},
	}

	for _, test := range tests {
		ctx := context.Background()

		actualValue, actualErr := ext.GetContextValues(ctx, test.Provided.ValuesKey)

		assert.Equal(t, test.Expected.Error.Error(), actualErr.Error(), test.Name)
		assert.Nil(t, actualValue, test.Name)
	}
}

func TestGetContextValuesKey(t *testing.T) {
	key := ext.ContextKey("key")
	value := "value"
	cMap := make(ext.ContextMap)
	cMap[key] = value

	type provided struct {
		ValuesKey  ext.ContextValuesKey
		ContextMap *ext.ContextMap
		Key        ext.ContextKey
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
				ValuesKey:  ext.ContextValuesKey("Values"),
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
		ctx = ext.CreateContextWithMap(ctx, test.Provided.ValuesKey, *test.Provided.ContextMap)

		actualValue, actualErr := ext.GetContextValuesKey(ctx, test.Provided.ValuesKey, test.Provided.Key)

		assert.Equal(t, test.Expected.Error, actualErr, test.Name)
		assert.Equal(t, test.Expected.Value, actualValue, test.Name)
	}
}

func TestGetContextValuesKeyError(t *testing.T) {
	emptyMap := make(ext.ContextMap)

	type provided struct {
		ValuesKey  ext.ContextValuesKey
		ContextMap *ext.ContextMap
		Key        ext.ContextKey
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
				ValuesKey:  ext.ContextValuesKey(""),
				ContextMap: &emptyMap,
				Key:        ext.ContextKey(""),
			},
			Expected: expected{
				Error: ext.ErrContextValuesNotFound,
				Value: nil,
			},
		},
		{
			Name: "Test missing Key error",
			Provided: provided{
				ValuesKey:  ext.ContextValuesKey("valuesKey"),
				ContextMap: &emptyMap,
				Key:        ext.ContextKey("key"),
			},
			Expected: expected{
				Error: ext.ErrContextValuesKeyNotFound,
				Value: nil,
			},
		},
	}

	for _, test := range tests {
		ctx := context.Background()
		if len(test.Provided.ValuesKey) > 0 {
			cm := make(ext.ContextMap)
			ctx = ext.CreateContextWithMap(ctx, test.Provided.ValuesKey, cm)
		}

		actualValue, actualErr := ext.GetContextValuesKey(ctx, test.Provided.ValuesKey, test.Provided.Key)

		assert.Equal(t, test.Expected.Error, actualErr, test.Name)
		assert.Nil(t, actualValue, test.Name)
	}
}
