package httpservice

import (
	"crypto-performance-compare/crypto"
	"crypto-performance-compare/utils"
	"encoding/json"
	"log"
	"net/http"
)

type HTTPHandler struct {
	cache *crypto.Cache
}

func newHandler(cache *crypto.Cache) HTTPHandler {
	return HTTPHandler{
		cache: cache,
	}
}

func (h *HTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res, err := h.cache.Read("BTC")
	if err != nil {
		http.Error(w, "no data found for BTC", http.StatusNotFound)
	}

	json.NewEncoder(w).Encode(res)
}

func NewServer(logger *log.Logger, cache *crypto.Cache) error {
	logger.Println("Starting HTTP server on http://localhost" + utils.GetPort())

	handler := newHandler(cache)
	if err := http.ListenAndServe(utils.GetPort(), &handler); err != nil {
		return err
	}

	return nil
}
