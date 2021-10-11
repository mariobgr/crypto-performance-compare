package crypto

import (
	"crypto-performance-compare/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	baseURL      string
	apiKey       string
	baseCurrency string
}

type Response struct {
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
	Price  string `json:"price"`
	Delta  string `json:"delta_1h"`
}

// NewClient returns *Client with config from env
func NewClient(baseURL string) *Client {
	return &Client{
		baseURL:      baseURL,
		apiKey:       utils.GetEnv("API_KEY", "default"),
		baseCurrency: utils.GetEnv("BASE_CURRENCY", "USD"),
	}
}

// GetInfo returns current stats for a coin from the API
func (c *Client) GetInfo(symbol string) (Response, error) {
	var response Response

	resp, err := http.Get(fmt.Sprintf("%s/api/v1/coin?key=%s&pref=%s&symbol=%s", c.baseURL, c.apiKey, c.baseCurrency, symbol))
	if err != nil {
		return response, fmt.Errorf("getting response from api: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return response, fmt.Errorf("bad response, status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, fmt.Errorf("reading response body: %w", err)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, fmt.Errorf("reading response body: %w", err)
	}

	return response, nil
}
