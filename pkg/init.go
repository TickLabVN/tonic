package pkg

import "github.com/TickLabVN/tonic/schema"

var globalSpec *schema.Root

func Init(options ...Config) {
	c := &schema.Root{
		OpenAPI: "3.0.0",
		Info: &schema.Info{
			Title:       "Simple API",
			Version:     "0.0.0",
			Description: "This is a simple API",
		},
		JsonSchemaDialect: "https://json-schema.org/draft/2020-12/schema",
		Servers: []*schema.Server{
			{
				URL:         "http://localhost:12345",
				Description: "Development server",
			},
		},
	}

	for _, option := range options {
		option.apply(c)
	}

	globalSpec = c
}
