package spec

import "github.com/TickLabVN/tonic/docs"

type Config interface {
	apply(*docs.OpenApi)
}

type configFn func(*docs.OpenApi)

func (fn configFn) apply(c *docs.OpenApi) {
	fn(c)
}

func WithOpenAPI(openAPI string) configFn {
	return func(c *docs.OpenApi) {
		c.OpenAPI = openAPI
	}
}

func WithInfo(info *docs.Info) configFn {
	return func(c *docs.OpenApi) {
		c.Info = info
	}
}

func WithServers(servers []*docs.Server) configFn {
	return func(c *docs.OpenApi) {
		c.Servers = servers
	}
}
