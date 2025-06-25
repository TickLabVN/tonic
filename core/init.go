package core

import "github.com/TickLabVN/tonic/core/docs"

func Init() *docs.OpenApi {
	c := &docs.OpenApi{
		OpenAPI: "3.0.0",
		Info: &docs.Infoobject{
			Title:   "Tonic API",
			Version: "0.0.0",
		},
	}

	return c
}
