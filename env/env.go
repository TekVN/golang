package env

import (
	"os"
	"strconv"
)

func GetEnvWithDefault(key string, defaultValue string) string {
	value := GetEnv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func GetEnv(key string) string {
	return os.Getenv(key)
}

func GetEnvInt(key string) (int, bool) {
	val, err := strconv.Atoi(GetEnv(key))
	if err != nil {
		return 0, false
	}
	return val, true
}

func GetEnvIntWithDefault(key string, defaultValue int) int {
	val, ok := GetEnvInt(key)
	if !ok {
		return defaultValue
	}
	return val
}
