package ext_test

import (
	"fmt"
	"testing"

	"github.com/mjdusa/go-ext"
	"github.com/stretchr/testify/assert"
)

func TestWrapError(t *testing.T) {
	err := fmt.Errorf("Test_WrapError: %d", 99)
	msg := "message"

	expected := fmt.Errorf("%s: %w", msg, err)

	actual := ext.WrapError(msg, err)

	assert.Equal(t, expected, actual)
}

func TestWrapErrorNoMessage(t *testing.T) {
	err := fmt.Errorf("Test_WrapError_noMessage: %d", 99)
	msg := ""

	expected := fmt.Errorf("%w", err)

	actual := ext.WrapError(msg, err)

	assert.Equal(t, expected, actual)
}
