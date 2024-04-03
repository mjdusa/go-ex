package envex

import (
	"os"
	"strconv"
	"strings"
	"time"
)

func SafeParseInt(str string, base int, bitSize int, defaultValue int64) int64 {
	result, err := strconv.ParseInt(str, base, bitSize)
	if err != nil {
		return defaultValue
	}
	return result
}

func GetEnvBool(envVar string, defaultValue bool) bool {
	val := os.Getenv(envVar)
	val = strings.TrimSpace(val)
	result, err := strconv.ParseBool(val)
	if err != nil {
		return defaultValue
	}
	return result
}

func GetEnvInt64(envVar string, base int, bitSize int, defaultValue int64) int64 {
	val := os.Getenv(envVar)
	val = strings.TrimSpace(val)
	result, err := strconv.ParseInt(val, base, bitSize)
	if err != nil {
		return defaultValue
	}
	return result
}

func GetEnvUInt64(envVar string, base int, bitSize int, defaultValue uint64) uint64 {
	val := os.Getenv(envVar)
	val = strings.TrimSpace(val)
	result, err := strconv.ParseUint(val, base, bitSize)
	if err != nil {
		return defaultValue
	}
	return result
}

func GetEnvString(envVar string, defaultValue string) string {
	val := os.Getenv(envVar)
	val = strings.TrimSpace(val)
	if len(val) == 0 {
		val = defaultValue
	}
	return val
}

func GetEnvDuration(envVar string, defaultValue time.Duration) time.Duration {
	val := os.Getenv(envVar)
	val = strings.TrimSpace(val)
	result, err := time.ParseDuration(val)
	if err != nil {
		return defaultValue
	}
	return result
}

func GetEnvTime(envVar string, layout string, defaultValue time.Time) time.Time {
	val := os.Getenv(envVar)
	val = strings.TrimSpace(val)
	result, err := time.Parse(layout, val)
	if err != nil {
		return defaultValue
	}
	return result
}
