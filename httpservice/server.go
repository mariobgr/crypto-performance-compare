package httpservice

import (
	"crypto-performance-compare/crypto"
	"crypto-performance-compare/utils"
	"html/template"
	"net/http"
)

type Server struct {
	*http.Server
}

type HTTPHandler struct {
	cache *crypto.Cache
}

func (h *HTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["Title"] = "Fake Titile"

	res, err := h.cache.Read("BTC")
	if err != nil {
		http.Error(w, "no data found for BTC", http.StatusNotFound)
	}

	data["Data"] = res

	tmpl, err := template.ParseFiles("./httpservice/template.html")
	if err != nil {
		http.Error(w, "error parsing template:" + err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "error serving data:" + err.Error(), http.StatusInternalServerError)
		return
	}
}

func NewServer(cache *crypto.Cache) Server {
	handler := newHandler(cache)

	return Server{
		&http.Server{
			Addr:    utils.GetPort(),
			Handler:  &handler,
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
