package util

import "os"

func GetEnv(key string, fallback string) string {
	if v, s := os.LookupEnv(key); s {
		return v
	}
	return fallback
}
