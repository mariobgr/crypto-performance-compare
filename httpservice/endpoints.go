package httpservice

import (
	"crypto-performance-compare/crypto"
	"crypto-performance-compare/utils"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func (h *HTTPHandler) ServeStats(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})

	tmpl, err := template.ParseFiles("./httpservice/template.html")
	if err != nil {
		http.Error(w, "error parsing template:"+err.Error(), http.StatusInternalServerError)
		return
	}

	items := utils.GetCurrencies()
	dataItems := make(map[string][]crypto.Response)

	for _, symbol := range items {
		data, err := h.cache.Read(symbol)
		if err != nil {
			http.Error(w, fmt.Sprintf("error reading data for %s : %s", symbol, err.Error()), http.StatusInternalServerError)
			return
		}
		dataItems[symbol] = data
	}

	data["title"] = fmt.Sprintf("Compare %s", strings.Join(items, ", "))
	data["dataItems"] = dataItems

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "error serving data:"+err.Error(), http.StatusInternalServerError)
		return
	}
}
