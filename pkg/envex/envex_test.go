package envex_test

import (
	"os"
	"testing"
	"time"

	"github.com/mjdusa/go-ext/pkg/envex"
	"github.com/stretchr/testify/assert"
)

func Test_SafeParseInt(t *testing.T) {
	assert := assert.New(t)

	testCases := []struct {
		IntString    string
		DefaultValue int64
		Expected     int64
		Description  string
	}{
		{
			IntString:    "",
			DefaultValue: int64(999),
			Expected:     int64(999),
			Description:  "empty string",
		},
		{
			IntString:    "-1",
			DefaultValue: int64(999),
			Expected:     int64(-1),
			Description:  "negative one",
		},
		{
			IntString:    "0",
			DefaultValue: int64(999),
			Expected:     int64(0),
			Description:  "zero",
		},
		{
			IntString:    "+1",
			DefaultValue: int64(999),
			Expected:     int64(1),
			Description:  "positive one",
		},
		{
			IntString:    "bad",
			DefaultValue: int64(999),
			Expected:     int64(999),
			Description:  "bad value, default",
		},
	}

	for _, tc := range testCases {
		actual := envex.SafeParseInt(tc.IntString, 10, 64, tc.DefaultValue)

		assert.Equal(tc.Expected, actual, tc.Description)
	}
}

func Test_GetEnvBool(t *testing.T) {
	assert := assert.New(t)
	envBadValue := "falsely"
	envGoodValue := "true"
	testCases := []struct {
		EnvKey       string
		EnvValue     *string
		DefaultValue bool
		Expected     bool
		Description  string
	}{
		{
			EnvKey:       "",
			EnvValue:     nil,
			DefaultValue: bool(false),
			Expected:     bool(false),
			Description:  "empty key",
		},
		{
			EnvKey:       "exampleKey",
			EnvValue:     nil,
			DefaultValue: bool(true),
			Expected:     bool(true),
			Description:  "nil env",
		},
		{
			EnvKey:       "exampleKey",
			EnvValue:     &envBadValue,
			DefaultValue: bool(true),
			Expected:     bool(true),
			Description:  "good env, bad value",
		},
		{
			EnvKey:       "exampleKey",
			EnvValue:     &envGoodValue,
			DefaultValue: bool(false),
			Expected:     bool(true),
			Description:  "good env, good value",
		},
	}

	for _, tc := range testCases {
		if tc.EnvValue == nil {
			os.Unsetenv(tc.EnvKey)
		} else {
			os.Setenv(tc.EnvKey, *tc.EnvValue)
		}

		actual := envex.GetEnvBool(tc.EnvKey, tc.DefaultValue)

		assert.Equal(tc.Expected, actual, tc.Description)
	}
}

func Test_GetEnvInt64(t *testing.T) {
	assert := assert.New(t)
	envBadValue := "bad"
	envGoodValue := "77777777777"
	testCases := []struct {
		EnvKey       string
		EnvValue     *string
		DefaultValue int64
		Expected     int64
		Description  string
	}{
		{
			EnvKey:       "",
			EnvValue:     nil,
			DefaultValue: int64(777),
			Expected:     int64(777),
			Description:  "empty key",
		},
		{
			EnvKey:       "exampleKey",
			EnvValue:     nil,
			DefaultValue: int64(777),
			Expected:     int64(777),
			Description:  "nil env",
		},
		{
			EnvKey:       "exampleKey",
			EnvValue:     &envBadValue,
			DefaultValue: int64(777),
			Expected:     int64(777),
			Description:  "good env, bad value",
		},
		{
			EnvKey:       "exampleKey",
			EnvValue:     &envGoodValue,
			DefaultValue: int64(777),
			Expected:     int64(77777777777),
			Description:  "good env, good value",
		},
	}

	for _, tc := range testCases {
		if tc.EnvValue == nil {
			os.Unsetenv(tc.EnvKey)
		} else {
			os.Setenv(tc.EnvKey, *tc.EnvValue)
		}

		actual := envex.GetEnvInt64(tc.EnvKey, 10, 64, tc.DefaultValue)

		assert.Equal(tc.Expected, actual, tc.Description)
	}
}

func Test_GetEnvUInt64(t *testing.T) {
	assert := assert.New(t)
	envBadValue := "bad"
	envGoodValue := "77777777777"
	testCases := []struct {
		EnvKey       string
		EnvValue     *string
		DefaultValue uint64
		Expected     uint64
		Description  string
	}{
		{
			EnvKey:       "",
			EnvValue:     nil,
			DefaultValue: uint64(777),
			Expected:     uint64(777),
			Description:  "empty key",
		},
		{
			EnvKey:       "exampleKey",
			EnvValue:     nil,
			DefaultValue: uint64(777),
			Expected:     uint64(777),
			Description:  "nil env",
		},
		{
			EnvKey:       "exampleKey",
			EnvValue:     &envBadValue,
			DefaultValue: uint64(777),
			Expected:     uint64(777),
			Description:  "good env, bad value",
		},
		{
			EnvKey:       "exampleKey",
			EnvValue:     &envGoodValue,
			DefaultValue: uint64(777),
			Expected:     uint64(77777777777),
			Description:  "good env, good value",
		},
	}

	for _, tc := range testCases {
		if tc.EnvValue == nil {
			os.Unsetenv(tc.EnvKey)
		} else {
			os.Setenv(tc.EnvKey, *tc.EnvValue)
		}

		actual := envex.GetEnvUInt64(tc.EnvKey, 10, 64, tc.DefaultValue)

		assert.Equal(tc.Expected, actual, tc.Description)
	}
}

func Test_GetEnvString(t *testing.T) {
	assert := assert.New(t)
	envValue := "envValue"
	testCases := []struct {
		EnvKey       string
		EnvValue     *string
		DefaultValue string
		Expected     string
		Description  string
	}{
		{
			EnvKey:       "",
			EnvValue:     nil,
			DefaultValue: "default",
			Expected:     "default",
			Description:  "empty key",
		},
		{
			EnvKey:       "exampleKey",
			EnvValue:     nil,
			DefaultValue: "default",
			Expected:     "default",
			Description:  "nil env",
		},
		{
			EnvKey:       "exampleKey",
			EnvValue:     &envValue,
			DefaultValue: "default",
			Expected:     envValue,
			Description:  "good env",
		},
	}

	for _, tc := range testCases {
		if tc.EnvValue == nil {
			os.Unsetenv(tc.EnvKey)
		} else {
			os.Setenv(tc.EnvKey, *tc.EnvValue)
		}

		actual := envex.GetEnvString(tc.EnvKey, tc.DefaultValue)

		assert.Equal(tc.Expected, actual, tc.Description)
	}
}

func Test_GetEnvDuration(t *testing.T) {
	assert := assert.New(t)
	envBadValue := "bad"
	envGoodValue := "777777ns"
	testCases := []struct {
		EnvKey       string
		EnvValue     *string
		DefaultValue time.Duration
		Expected     time.Duration
		Description  string
	}{
		{
			EnvKey:       "",
			EnvValue:     nil,
			DefaultValue: time.Duration(777),
			Expected:     time.Duration(777),
			Description:  "empty key",
		},
		{
			EnvKey:       "exampleKey",
			EnvValue:     nil,
			DefaultValue: time.Duration(777),
			Expected:     time.Duration(777),
			Description:  "nil env",
		},
		{
			EnvKey:       "exampleKey",
			EnvValue:     &envBadValue,
			DefaultValue: time.Duration(777),
			Expected:     time.Duration(777),
			Description:  "good env, bad value",
		},
		{
			EnvKey:       "exampleKey",
			EnvValue:     &envGoodValue,
			DefaultValue: time.Duration(777),
			Expected:     time.Duration(777777 * time.Nanosecond),
			Description:  "good env, good value",
		},
	}

	for _, tc := range testCases {
		if tc.EnvValue == nil {
			os.Unsetenv(tc.EnvKey)
		} else {
			os.Setenv(tc.EnvKey, *tc.EnvValue)
		}

		actual := envex.GetEnvDuration(tc.EnvKey, tc.DefaultValue)

		assert.Equal(tc.Expected, actual, tc.Description)
	}
}

func Test_GetEnvTime(t *testing.T) {
	assert := assert.New(t)
	layout := time.RFC3339
	envBadValue := "bad"
	envGoodValue := "2023-01-02T03:04:05-06:00"
	expected, err := time.Parse(layout, envGoodValue)
	if err != nil {
		assert.FailNow("Test setup error: %w", err)
	}

	now := time.Now()
	testCases := []struct {
		Layout       string
		EnvKey       string
		EnvValue     *string
		DefaultValue time.Time
		Expected     time.Time
		Description  string
	}{
		{
			Layout:       layout,
			EnvKey:       "",
			EnvValue:     nil,
			DefaultValue: now,
			Expected:     now,
			Description:  "empty key",
		},
		{
			Layout:       layout,
			EnvKey:       "exampleKey",
			EnvValue:     nil,
			DefaultValue: now,
			Expected:     now,
			Description:  "nil env",
		},
		{
			Layout:       layout,
			EnvKey:       "exampleKey",
			EnvValue:     &envBadValue,
			DefaultValue: now,
			Expected:     now,
			Description:  "good env, bad value",
		},
		{
			Layout:       "",
			EnvKey:       "exampleKey",
			EnvValue:     &envGoodValue,
			DefaultValue: now,
			Expected:     now,
			Description:  "good env, bad layout, good value",
		},
		{
			Layout:       layout,
			EnvKey:       "exampleKey",
			EnvValue:     &envGoodValue,
			DefaultValue: now,
			Expected:     expected,
			Description:  "good env, good layout, good value",
		},
	}

	for _, tc := range testCases {
		if tc.EnvValue == nil {
			os.Unsetenv(tc.EnvKey)
		} else {
			os.Setenv(tc.EnvKey, *tc.EnvValue)
		}

		actual := envex.GetEnvTime(tc.EnvKey, tc.Layout, tc.DefaultValue)

		assert.Equal(tc.Expected, actual, tc.Description)
	}
}
