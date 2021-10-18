package utils

import (
	"log"
	"os"
)

const (
	ColorSuccess = "\033[32m"
	ColorError   = "\033[31m"
	ColorReset   = "\033[0m"
)

func NewLogger() *log.Logger {
	logger := log.New(os.Stdout, "", log.Ldate | log.Ltime)
	return logger
}
