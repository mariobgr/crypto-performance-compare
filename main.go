package main

import (
	"crypto-performance-compare/crypto"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"time"
)

var (
	colorSuccess = "\033[32m"
	colorError   = "\033[31m"
)

func main() {
	start := time.Now()

	// Load env vars
	godotenv.Load()

	// Init Crypto Client
	client := crypto.NewClient(os.Getenv("BASE_URL"))

	// TODO: Read list of tracked coins from .env
	_, err := client.GetInfo("BTC")
	if err != nil {
		fmt.Println(colorError, "error getting info for BTC:", err.Error())
		return
	}

	fmt.Println(colorSuccess, "Successfully init app in", time.Since(start))
}
