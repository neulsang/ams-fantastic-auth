package env

import (
	"os"
	"strconv"
)

func ReadAsStr(key string, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}

	return defaultVal
}

func ReadAsInt(key string, defaultVal int) int {
	strVal := ReadAsStr(key, "")

	if val, err := strconv.Atoi(strVal); err == nil {
		return val
	}
	return defaultVal
}
