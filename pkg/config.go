package pkg

import "github.com/TickLabVN/tonic/schema"

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
