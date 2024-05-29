package pkg

import "github.com/TickLabVN/tonic/docs"

var globalSpec *docs.OpenApi

func Init(options ...Config) {
	c := &docs.OpenApi{
		OpenAPI: "3.0.0",
		Info: &docs.Info{
			Title:       "Simple API",
			Version:     "0.0.0",
			Description: "This is a simple API",
		},
		JsonSchemaDialect: "https://json-schema.org/draft/2020-12/schema",
	}

	for _, option := range options {
		option.apply(c)
	}

	globalSpec = c
}
