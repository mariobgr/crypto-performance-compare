package main

import (
	"crypto-performance-compare/crypto"
	"crypto-performance-compare/utils"
	"github.com/joho/godotenv"
	"time"
)

func main() {
	start := time.Now()

	// Load env vars
	godotenv.Load()

	// Init custom logger
	logger := utils.NewLogger()

	// Init cache
	cache := crypto.NewCache()

	// Init Crypto Updater
	updater := crypto.NewUpdater(logger, cache)

	// Start the continuous update procedure
	err := updater.UpdateAll()
	if err != nil {
		logger.Println(utils.ColorError, "Error starting the update procedure:", err.Error())
		return
	}

	// TODO: Make service run indefinitely in background
	time.Sleep(10 * time.Second)

	// TODO: Remove debug
	logger.Println(cache.Read("BTC"))
	logger.Println(cache.Read("ETH"))
	logger.Println(cache.Read("SHIB"))

	logger.Println(utils.ColorSuccess, "Successfully init app in", time.Since(start))
}
