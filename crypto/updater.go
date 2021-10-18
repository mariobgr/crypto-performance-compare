package crypto

import (
	"crypto-performance-compare/utils"
	"fmt"
	"log"
	"os"
	"time"
)

type Updater struct {
	logger *log.Logger
	client *Client
}

func NewUpdater(logger *log.Logger, cache *Cache) *Updater {
	return &Updater{
		logger: logger,
		client: NewClient(os.Getenv("BASE_URL"), cache),
	}
}

func (u *Updater) UpdateAll() error {
	items := utils.GetCurrencies()

	if len(items) == 0 {
		return fmt.Errorf("no currencies selected - please set them up as a csv in ENV variables")
	}

	// Eagerly update all
	for _, item := range items {
		u.logger.Println("Starting update for", item)
		u.Update(item)
	}

	// Set up update procedure every minute
	ticker := time.NewTicker(time.Minute)
	go func() {
		for range ticker.C {
			for _, item := range items {
				u.logger.Println("Start update for", item)
				u.Update(item)
			}
		}
	}()

	return nil
}

func (u *Updater) Update(symbol string) {
	res, err := u.client.GetInfo(symbol)
	if err != nil {
		u.logger.Println(utils.ColorError, "error getting info for BTC:", err.Error(), utils.ColorReset)
		return
	}

	u.client.Cache.Add(symbol, res)

	u.logger.Println(utils.ColorSuccess, "Successfully updated cache for", symbol, utils.ColorReset)
}
