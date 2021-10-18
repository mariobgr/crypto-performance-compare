package utils

import (
	"os"
	"strings"
)

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func GetCurrencies() []string {
	items := GetEnv("TRACK_LIST", "")
	return strings.Split(items, ",")
}

func GetPort() string {
	return GetEnv("SERVER_PORT", ":8080")
}
