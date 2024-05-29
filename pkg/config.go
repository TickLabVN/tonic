package pkg

import "github.com/TickLabVN/tonic/schema"

type Config interface {
	apply(*schema.OpenApi)
}

type configFn func(*schema.OpenApi)

func (fn configFn) apply(c *schema.OpenApi) {
	fn(c)
}

func WithOpenAPI(openAPI string) configFn {
	return func(c *schema.OpenApi) {
		c.OpenAPI = openAPI
	}
}

func WithInfo(info *schema.Info) configFn {
	return func(c *schema.OpenApi) {
		c.Info = info
	}
}

func WithServers(servers []interface{}) configFn {
	return func(c *schema.OpenApi) {
		c.Servers = servers
	}
}
