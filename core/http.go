package core

import (
	_ "embed"
	"encoding/json"
	"net/http"

	"github.com/TickLabVN/tonic/core/docs"
	"github.com/swaggo/http-swagger"
)

func JsonHttpHandler(spec *docs.OpenApi) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(spec)
	})
}

func SwaggerUIHandler(url string) http.Handler {
	return httpSwagger.Handler(
		httpSwagger.URL(url),
	)
}
