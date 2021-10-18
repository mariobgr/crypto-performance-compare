package crypto_test

import (
	"crypto-performance-compare/crypto"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_GetInfoSuccess(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, successResponse)
	}))

	client := crypto.NewClient(ts.URL)
	res, err := client.GetInfo("BTC")
	require.NoError(t, err)
	assert.NotNil(t, res.Name)
	assert.NotNil(t, res.Price)
}

func TestClient_GetInfoFailure(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"success": false}`)
	}))

	client := crypto.NewClient(ts.URL)
	res, err := client.GetInfo("BTC")
	require.Error(t, err)
	assert.Nil(t, res.Name)
	assert.Nil(t, res.Price)
}

var successResponse = `{
	"symbol": "BTC",
	"show_symbol": "BTC",
	"name": "Bitcoin",
	"rank": 1,
	"price": "5524.7112165586",
	"market_cap": "94433817003.39",
	"total_volume_24h": "6378793658.5432",
	"low_24h": "5324.2665427149",
	"high_24h": "5561.0068476948",
	"delta_1h": "0.81",
	"delta_24h": "0.68",
	"delta_7d": "-15.26",
	"delta_30d": "-25.26",
	"markets": [
		{
			"symbol": "EUR",
			"volume_24h": "123707000",
			"price": "5524.7112165586",
			"exchanges": [
				{
					"name": "Kraken",
					"volume_24h": "50623900",
					"price": "5520"
				},
				{
					"name": "Bitfinex",
					"volume_24h": "19314700",
					"price": "5512.6"
				}
			]
		}
	],
	"last_updated_timestamp": 1528987416,
	"remaining": 1133
}`
