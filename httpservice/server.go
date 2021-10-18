package httpservice

import (
	"crypto-performance-compare/crypto"
	"crypto-performance-compare/utils"
	"net/http"
)

type Server struct {
	*http.Server
}

type HTTPHandler struct {
	cache *crypto.Cache
}

func (h *HTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.ServeStats(w, r)
}

func NewServer(cache *crypto.Cache) Server {
	handler := newHandler(cache)

	return Server{
		&http.Server{
			Addr:    utils.GetPort(),
			Handler: &handler,
		},
	}
}

func (s Server) Start() error {
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}
	return nil
}

func newHandler(cache *crypto.Cache) HTTPHandler {
	return HTTPHandler{
		cache: cache,
	}
}
