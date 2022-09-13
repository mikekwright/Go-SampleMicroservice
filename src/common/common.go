package common

import "os"

func GetEnvOrDefault(env string, fallback string) string {
	if value, exists := os.LookupEnv(env); exists {
		return value
	} else {
		return fallback
	}
}
