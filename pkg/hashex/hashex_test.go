package hashex_test

import (
	"testing"

	"github.com/mjdusa/go-ext/pkg/hashex"
	"github.com/stretchr/testify/assert"
)

func TestSHA256(t *testing.T) {
	type provided struct {
		Data string
	}
	type expected struct {
		Error error
		Value string
	}

	tests := []struct {
		Name     string
		Provided provided
		Expected expected
	}{
		{
			Name: "SHA256 test",
			Provided: provided{
				Data: "abcdefghijklmnipqrstuvwxyz01234567890`-=_+[]{}|\\:;'\",.<>?/~!@#$%^&*()_+",
			},
			Expected: expected{
				Error: nil,
				Value: "74ac7f49882655aa1abfe7933011c70214ea21630476c8b5b2a8287c9be7cf59",
			},
		},
	}

	for _, test := range tests {
		actual := hashex.SHA256(test.Provided.Data)

		assert.Equal(t, test.Expected.Value, actual, test.Name)
	}
}

func TestSHA512(t *testing.T) {
	type provided struct {
		Data string
	}
	type expected struct {
		Error error
		Value string
	}

	tests := []struct {
		Name     string
		Provided provided
		Expected expected
	}{
		{
			Name: "SHA512 test",
			Provided: provided{
				Data: "abcdefghijklmnipqrstuvwxyz01234567890`-=_+[]{}|\\:;'\",.<>?/~!@#$%^&*()_+",
			},
			Expected: expected{
				Error: nil,
				//nolint:lll,nolintlint  // test string
				Value: "97e8ff148518116eb7b2e023c4db798c3b854bc574b712ee941671dd16d57634c6a51ae6c49f6d46c8a939fef767a2e116e4247c6f4bdb4265e2258fb9835c83",
			},
		},
	}

	for _, test := range tests {
		actual := hashex.SHA512(test.Provided.Data)

		assert.Equal(t, test.Expected.Value, actual, test.Name)
	}
}
