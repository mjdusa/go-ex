package mapex_test

import (
	"reflect"
	"testing"

	"github.com/mjdusa/go-ext/pkg/mapex"
	"github.com/stretchr/testify/assert"
)

func TestPair(t *testing.T) {
	tests := []struct {
		name string
		args []interface{}
		want mapex.SIMap
	}{
		{
			name: "Test with no arguments",
			args: []interface{}{},
			want: make(mapex.SIMap),
		},
		{
			name: "Test with one pair",
			args: []interface{}{mapex.SIKey("key1"), "value1"},
			want: mapex.SIMap{mapex.SIKey("key1"): "value1"},
		},
		{
			name: "Test with multiple pairs",
			args: []interface{}{mapex.SIKey("key1"), "value1", mapex.SIKey("key2"), "value2"},
			want: mapex.SIMap{mapex.SIKey("key1"): "value1", mapex.SIKey("key2"): "value2"},
		},
		{
			name: "Test with nested pairs",
			args: []interface{}{mapex.SIKey("key1"), []interface{}{mapex.SIKey("nestedKey"), "nestedValue"}},
			want: mapex.SIMap{mapex.SIKey("key1"): mapex.SIMap{mapex.SIKey("nestedKey"): "nestedValue"}},
		},
		{
			name: "Test with one pair, bad key",
			args: []interface{}{999, "value1"},
			want: mapex.SIMap{},
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := mapex.Pair(tst.args...)

			if !reflect.DeepEqual(got, tst.want) {
				assert.Fail(t, "Got of %v not equal to want %v", got, tst.want)
			}
		})
	}
}
