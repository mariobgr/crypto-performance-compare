package crypto_test

import (
	"crypto-performance-compare/crypto"
	"crypto-performance-compare/fakes"
	"crypto-performance-compare/utils"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestUpdater(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, fakes.SuccessResponse)
	}))

	os.Setenv("BASE_URL", ts.URL)
	os.Setenv("TRACK_LIST", "BTC")

	logger := utils.NewLogger()
	cache := crypto.NewCache()

	updater := crypto.NewUpdater(logger, cache)
	err := updater.UpdateAll()
	require.NoError(t, err)

	items := utils.GetCurrencies()
	for _, item := range items {
		res, err := cache.Read(item)
		require.NoError(t, err)
		assert.NotEmpty(t, res)
	}

	res2, err := cache.Read("ETH")
	assert.Error(t, err)
	assert.Empty(t, res2)
}