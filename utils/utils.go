package utils

import (
	"fmt"
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
		return fmt.Sprintf("%.2f B", float64(bytes))
	} else if bytes < (1 << 20) {
		return fmt.Sprintf("%.2f KB", float64(bytes)/(1<<10))
	} else if bytes < (1 << 30) {
		return fmt.Sprintf("%.2f MB", float64(bytes)/(1<<20))
	} else {
		return fmt.Sprintf("%.2f GB", float64(bytes)/(1<<30))
	}
}
