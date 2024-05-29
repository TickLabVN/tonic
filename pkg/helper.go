package pkg

import "github.com/TickLabVN/tonic/docs"

func GetSpec() *docs.OpenApi {
	if globalSpec == nil {
		panic("spec is not initialized")
	}

	return globalSpec
}

func SetPath(path map[string]*docs.Path) {
	globalSpec.Paths = path
}
