package pkg

import "github.com/TickLabVN/tonic/schema"

var Spec *schema.Root

type Config interface {
	apply(*schema.Root)
}

type configFn func(*schema.Root)

func (fn configFn) apply(c *schema.Root) {
	fn(c)
}

func WithOpenAPI(openAPI string) configFn {
	return func(c *schema.Root) {
		c.OpenAPI = openAPI
	}
}

func WithInfo(info *schema.Info) configFn {
	return func(c *schema.Root) {
		c.Info = info
	}
}

func WithServers(servers []*schema.Server) configFn {
	return func(c *schema.Root) {
		c.Servers = servers
	}
}

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

	Spec = c
}

func GetSpec() *schema.Root {
	if Spec == nil {
		panic("spec is not initialized")
	}

	return Spec
}
