package core

import "github.com/TickLabVN/tonic/core/docs"

var globalSpec *docs.OpenApi

func Init(options ...Config) {
	c := &docs.OpenApi{
		OpenAPI: "3.0.0",
		Info: &docs.Info{
			Title:       "Tonic API",
			Version:     "0.0.0",
		},
	}

	for _, option := range options {
		option.apply(c)
	}

	globalSpec = c
}
