package tonic

import (
	"encoding/json"
	"net/http"

	"github.com/flowchartsman/swaggerui"
)

type ExternalDocs struct {
	Description string `json:"description"`
	URL         string `json:"url"`
}

type Config struct {
	OpenAPIVersion string         `json:"openapi"`
	Info           map[string]any `json:"info"`
	ExternalDocs   *ExternalDocs  `json:"externalDocs,omitempty"`
}

var isInit = false

func Init(config *Config) {
	apiSpec = make(map[string]any)

	apiSpec["openapi"] = config.OpenAPIVersion
	if config.Info != nil {
		apiSpec["info"] = config.Info
	}
	if config.ExternalDocs != nil {
		apiSpec["externalDocs"] = config.ExternalDocs
	}
	apiSpec["components"] = map[string]any{
		"schemas": make(map[string]any),
	}
	apiSpec["paths"] = make(map[string]any)
	isInit = true
}

func GetHandler() http.Handler {
	bytes, _ := json.Marshal(apiSpec)
	return swaggerui.Handler(bytes)
}
