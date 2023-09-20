package ext

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Reverse(t *testing.T) {
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
			Name: "Test Reverse empty",
			Provided: provided{
				Data: "",
			},
			Expected: expected{
				Error: nil,
				Value: "",
			},
		},
		{
			Name: "Test Reverse string",
			Provided: provided{
				Data: "123456",
			},
			Expected: expected{
				Error: nil,
				Value: "654321",
			},
		},
	}

	for _, test := range tests {
		actual := Reverse(test.Provided.Data)

		assert.Equal(t, test.Expected.Value, actual, test.Name)
	}
}

func Test_Ellipse(t *testing.T) {
	type provided struct {
		Data    string
		Max     int
		Ellipse string
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
			Name: "Test Ellipse empty string",
			Provided: provided{
				Data:    "",
				Max:     0,
				Ellipse: "...",
			},
			Expected: expected{
				Error: nil,
				Value: "",
			},
		},
		{
			Name: "Test Ellipse empty ellipse",
			Provided: provided{
				Data:    "你好 hello",
				Max:     6,
				Ellipse: "",
			},
			Expected: expected{
				Error: nil,
				Value: "你好 hel",
			},
		},
		{
			Name: "Test Ellipse normal no change",
			Provided: provided{
				Data:    "你好 hello",
				Max:     8,
				Ellipse: "...",
			},
			Expected: expected{
				Error: nil,
				Value: "你好 hello",
			},
		},
		{
			Name: "Test Ellipse normal ellipsed",
			Provided: provided{
				Data:    "你好 hello",
				Max:     6,
				Ellipse: "...",
			},
			Expected: expected{
				Error: nil,
				Value: "你好 ...",
			},
		},
	}

	for _, test := range tests {
		actual := Ellipse(test.Provided.Data, test.Provided.Max, test.Provided.Ellipse)

		assert.Equal(t, test.Expected.Value, actual, test.Name)
	}
}
