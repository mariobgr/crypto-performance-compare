package crypto_test

import (
	"crypto-performance-compare/crypto"
	"crypto-performance-compare/fakes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_GetInfoSuccess(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, fakes.SuccessResponse)
	}))

	client := crypto.NewClient(ts.URL, crypto.NewCache())
	res, err := client.GetInfo("BTC")

	require.NoError(t, err)
	assert.NotEmpty(t, res.Name)
	assert.NotEmpty(t, res.Price)
}

func TestClient_GetInfoFailure(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"success": false}`)
	}))

	client := crypto.NewClient(ts.URL, crypto.NewCache())
	res, err := client.GetInfo("BTC")

	require.Error(t, err)
	assert.Empty(t, res.Name)
	assert.Empty(t, res.Price)
}
