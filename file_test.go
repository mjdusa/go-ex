package ext

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FileSize(t *testing.T) {
	type provided struct {
		Data string
	}
	type expected struct {
		Error error
		Value *int64
	}

	size := int64(23)

	tests := []struct {
		Name     string
		Provided provided
		Expected expected
	}{
		{
			Name: "FileSize test - good file",
			Provided: provided{
				Data: "tests/text-test.data",
			},
			Expected: expected{
				Error: nil,
				Value: &size,
			},
		},
		{
			Name: "FileSize test - bad file",
			Provided: provided{
				Data: "tests/i-do-not-exist",
			},
			Expected: expected{
				Error: fmt.Errorf("os.Stat() error: stat tests/i-do-not-exist: no such file or directory"),
				Value: nil,
			},
		},
	}

	for _, test := range tests {
		actualValue, actualError := FileSize(test.Provided.Data)

		if actualError != nil {
			assert.Equal(t, test.Expected.Error.Error(), actualError.Error(), test.Name)
		}

		if actualValue != nil {
			assert.Equal(t, *test.Expected.Value, *actualValue, test.Name)
		}
	}
}
