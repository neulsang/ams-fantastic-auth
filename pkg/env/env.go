package env

import (
	"os"
	"strconv"
)

func Get(key string, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}

	return defaultVal
}

func GetAsInt(key string, defaultVal int) int {
	strVal := Get(key, "")

	if val, err := strconv.Atoi(strVal); err == nil {
		return val
	}
	return defaultVal
}
