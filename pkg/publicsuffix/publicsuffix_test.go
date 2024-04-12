package publicsuffix_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/mjdusa/go-ex/pkg/publicsuffix"
	"github.com/stretchr/testify/assert"
)

var (
	normalTimeout = 5 * time.Second         //nolint:gochecknoglobals  // Test time.Duration used by multiple tests
	shortTimeout  = 1 * time.Nanosecond     //nolint:gochecknoglobals  // Test time.Duration used by multiple tests
	cacheFileName = "../../tests/tld.cache" //nolint:gochecknoglobals  // Test file path used by multiple tests
)

func TestWrapError(t *testing.T) {
	err := fmt.Errorf("TestWrapError: %d", 99)
	msg := "message"

	expected := fmt.Errorf("%s: %w", msg, err)

	actual := publicsuffix.WrapError(msg, err)

	assert.Equal(t, expected, actual)
}

func TestLoadTLDsTimeout(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	urls := []string{publicsuffix.MainPublicSuffixListFile, publicsuffix.BackupPublicSuffixListFile}
	emptyFileName := ""
	useASCII := true

	actualList, actualErr := publicsuffix.LoadTLDs(ctx, urls, shortTimeout, emptyFileName, useASCII)

	assert.NotNil(actualErr, "Error Not Nil")
	assert.Nil(actualList, "List Nil")
}

func TestLoadTLDsCacheFile(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	urls := []string{}
	useASCII := true

	actualList, actualErr := publicsuffix.LoadTLDs(ctx, urls, normalTimeout, cacheFileName, useASCII)

	assert.Nil(actualErr, "Error Nil")
	assert.NotNil(actualList, "List Not Nil")
	assert.NotEqual(0, len(actualList), "List length > 0")
}

func TestLoadTLDsASCII(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	urls := []string{publicsuffix.MainPublicSuffixListFile}
	useASCII := true

	actualList, actualErr := publicsuffix.LoadTLDs(ctx, urls, normalTimeout, cacheFileName, useASCII)

	assert.Nil(actualErr, "Error Nil")
	assert.NotNil(actualList, "List Not Nil")
	assert.NotEqual(0, len(actualList), "List length > 0")
}

func TestLoadTLDsUnicode(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	urls := []string{publicsuffix.BackupPublicSuffixListFile}
	useASCII := false

	actualList, actualErr := publicsuffix.LoadTLDs(ctx, urls, normalTimeout, cacheFileName, useASCII)

	assert.Nil(actualErr, "Error Nil")
	assert.NotNil(actualList, "List Not Nil")
	assert.NotEqual(0, len(actualList), "List length > 0")
}
