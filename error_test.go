package ext

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_WrapError(t *testing.T) {
	err := fmt.Errorf("Test_WrapError: %d", 99)
	msg := "message"

	expected := fmt.Errorf("%s: %w", msg, err)

	actual := WrapError(msg, err)

	assert.Equal(t, expected, actual)
}
