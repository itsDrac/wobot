package utils

import (
	"os"
	"strconv"
)

func GetStringEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func GetIntEnv(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return intValue
}

func FormatBytes(bytes int64) string {
	if bytes < 1024 {
		return strconv.FormatInt(bytes, 10) + " B"
	} else if bytes < (1 << 20) {
		return strconv.FormatInt(bytes/(1<<10), 10) + " KB"
	} else if bytes < (1 << 30) {
		return strconv.FormatInt(bytes/(1<<20), 10) + " MB"
	} else {
		return strconv.FormatInt(bytes/(1<<30), 10) + " GB"
	}
}
