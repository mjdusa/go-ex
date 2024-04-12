package fileex_test

import (
	"fmt"
	"testing"

	"github.com/mjdusa/go-ex/pkg/fileex"
	"github.com/stretchr/testify/assert"
)

func TestWrapError(t *testing.T) {
	err := fmt.Errorf("TestWrapError: %d", 99)
	msg := "message"

	expected := fmt.Errorf("%s: %w", msg, err)

	actual := fileex.WrapError(msg, err)

	assert.Equal(t, expected, actual)
}

func TestFileSize(t *testing.T) {
	existingFile := "../../tests/text-test.data"
	missingFile := "../../tests/i-do-not-exist"

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
				Data: existingFile,
			},
			Expected: expected{
				Error: nil,
				Value: &size,
			},
		},
		{
			Name: "FileSize test - bad file",
			Provided: provided{
				Data: missingFile,
			},
			Expected: expected{
				Error: fmt.Errorf("os.Stat() error: stat %s: no such file or directory", missingFile),
				Value: nil,
			},
		},
	}

	for _, test := range tests {
		actualValue, actualError := fileex.FileSize(test.Provided.Data)

		if actualError != nil {
			assert.Equal(t, test.Expected.Error.Error(), actualError.Error(), test.Name)
		}

		if actualValue != nil {
			assert.Equal(t, *test.Expected.Value, *actualValue, test.Name)
		}
	}
}
