package env

import (
	"fmt"
	"os"
	"strconv"
)

func String(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}

func Int(key string, fallback int) int {
	if val, ok := os.LookupEnv(key); ok {
		i, err := strconv.Atoi(val)
		if err != nil {
			panic(fmt.Sprintf("environment variable '%s' is not an integer", key))
		}
		return i
	}
	return fallback
}

func MustString(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic(fmt.Sprintf("environment variable '%s' is required", key))
	}
	return val
}

func MustInt(key string) int {
	val, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		panic(fmt.Sprintf("environment variable '%s' is not an integer", key))
	}
	return val
}
