package env

import (
	"fmt"
	"os"
)

func GetEnvOrDefault(env, defaultValue string) string {
	value, ok := os.LookupEnv(env)
	if !ok {
		value = defaultValue
	}
	return value
}

func GetEnvOrError(env string) (string, error) {
	value, ok := os.LookupEnv(env)
	if !ok {
		return "", fmt.Errorf("%s is not set", env)
	}
	return value, nil
}
