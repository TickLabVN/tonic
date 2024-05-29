package pkg

import "github.com/TickLabVN/tonic/schema"

func GetSpec() *schema.OpenApi {
	if globalSpec == nil {
		panic("spec is not initialized")
	}

	return globalSpec
}

func SetPath(path map[string]*schema.Path) {
	globalSpec.Paths = path
}
