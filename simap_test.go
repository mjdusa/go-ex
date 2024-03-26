package ext_test

import (
	"reflect"
	"testing"

	"github.com/mjdusa/go-ext"
	"github.com/stretchr/testify/assert"
)

func TestPair(t *testing.T) {
	tests := []struct {
		name string
		args []interface{}
		want ext.SIMap
	}{
		{
			name: "Test with no arguments",
			args: []interface{}{},
			want: make(ext.SIMap),
		},
		{
			name: "Test with one pair",
			args: []interface{}{ext.SIKey("key1"), "value1"},
			want: ext.SIMap{ext.SIKey("key1"): "value1"},
		},
		{
			name: "Test with multiple pairs",
			args: []interface{}{ext.SIKey("key1"), "value1", ext.SIKey("key2"), "value2"},
			want: ext.SIMap{ext.SIKey("key1"): "value1", ext.SIKey("key2"): "value2"},
		},
		{
			name: "Test with nested pairs",
			args: []interface{}{ext.SIKey("key1"), []interface{}{ext.SIKey("nestedKey"), "nestedValue"}},
			want: ext.SIMap{ext.SIKey("key1"): ext.SIMap{ext.SIKey("nestedKey"): "nestedValue"}},
		},
		{
			name: "Test with one pair, bad key",
			args: []interface{}{999, "value1"},
			want: ext.SIMap{},
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := ext.Pair(tst.args...)

			if !reflect.DeepEqual(got, tst.want) {
				assert.Fail(t, "Got of %v not equal to want %v", got, tst.want)
			}
		})
	}
}
