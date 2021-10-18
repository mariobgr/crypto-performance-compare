package crypto_test

import (
	"crypto-performance-compare/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	cache := crypto.NewCache()
	symbol := "BTC"
	now := time.Now()

	response := crypto.Response{
		Symbol: symbol,
		Name:   "Bitcoin",
		Price:  "60000",
		Delta:  "15%",
		Time:   now.Format("2006-01-02T15:04:05"),
	}

	cache.Add(symbol, response)

	result, err := cache.Read(symbol)
	require.NoError(t, err)
	assert.NotEmpty(t, result)

	result2, err := cache.Read("ETH")
	require.Error(t, err)
	assert.Empty(t, result2)
}
