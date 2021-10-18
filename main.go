package main

import (
	"crypto-performance-compare/crypto"
	"crypto-performance-compare/httpservice"
	"crypto-performance-compare/utils"
	"github.com/joho/godotenv"
)

func main() {
	// Load env vars
	godotenv.Load()

	// Init custom logger
	logger := utils.NewLogger()
	logger.Println("Starting service....")

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

	// Init & start HTTP server
	srv := httpservice.NewServer(cache)

	logger.Println("Starting HTTP server on http://localhost" + utils.GetPort())
	err = srv.Start()
	if err != nil {
		logger.Println(utils.ColorError, "Error starting HTTP server:", err.Error())
		return
	}
}
