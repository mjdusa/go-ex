package stringex_test

import (
	"testing"

	"github.com/mjdusa/go-ex/pkg/stringex"
	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
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
		actual := stringex.Reverse(test.Provided.Data)

		assert.Equal(t, test.Expected.Value, actual, test.Name)
	}
}

func TestEllipse(t *testing.T) {
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
		actual := stringex.Ellipse(test.Provided.Data, test.Provided.Max, test.Provided.Ellipse)

		assert.Equal(t, test.Expected.Value, actual, test.Name)
	}
}

func TestStringArrayContains(t *testing.T) {
	assert := assert.New(t)

	testCases := []struct {
		StringArray     []string
		Value           string
		CaseInsensitive bool
		Expected        bool
		Description     string
	}{
		{
			StringArray:     []string{},
			Value:           "",
			CaseInsensitive: false,
			Expected:        false,
			Description:     "Empty[] and empty Vaslue should return false",
		},
		{
			StringArray:     []string{},
			Value:           "Foo",
			CaseInsensitive: false,
			Expected:        false,
			Description:     "Empty[] with Value should return false",
		},
		{
			StringArray:     []string{"Foo"},
			Value:           "Foo",
			CaseInsensitive: false,
			Expected:        true,
			Description:     "Foo == Foo should return true",
		},
		{
			StringArray:     []string{"foo"},
			Value:           "Foo",
			CaseInsensitive: false,
			Expected:        false,
			Description:     "foo != Foo should return false",
		},
		{
			StringArray:     []string{"foo"},
			Value:           "Foo",
			CaseInsensitive: true,
			Expected:        true,
			Description:     "foo == Foo should return true",
		},
		{
			StringArray:     []string{"fee", "fi", "foe", "fum"},
			Value:           "Fi",
			CaseInsensitive: true,
			Expected:        true,
			Description:     "fi == Fi should return true",
		},
	}

	for _, tc := range testCases {
		actual := stringex.StringArrayContains(tc.StringArray, tc.Value, tc.CaseInsensitive)
		assert.Equal(tc.Expected, actual, tc.Description)
	}
}
