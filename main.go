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
	symbol := "BTC"
	res, err := client.GetInfo(symbol)
	if err != nil {
		fmt.Println(colorError, "error getting info for BTC:", err.Error())
		return
	}

	symbol2 := "ETH"
	res2, err := client.GetInfo(symbol2)
	if err != nil {
		fmt.Println(colorError, "error getting info for BTC:", err.Error())
		return
	}

	client.Cache.Add(symbol, res)
	client.Cache.Add(symbol2, res2)

	fmt.Println(client.Cache.Read(symbol))
	fmt.Println(client.Cache.Read(symbol2))

	fmt.Println(colorSuccess, "Successfully init app in", time.Since(start))
}
