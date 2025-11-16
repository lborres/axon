package config

import (
	"log"
	"os"
	"strconv"
)

func getEnvStr(key, fallback string, req bool) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	} else if !ok && req {
		log.Fatalf("Env key, %s, does not have a value", key)
	}
	return fallback
}

func getEnvInt64(key string, fallback int64, req bool) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			if req {
				log.Fatalf("Env key, %s, does not have a value", key)
			}
			return fallback
		}
		return i
	}
	if req {
		log.Fatalf("Env key, %s, does not have a value", key)
	}
	return fallback
}
