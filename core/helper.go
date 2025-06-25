package core

import "github.com/TickLabVN/tonic/core/docs"

func GetSpec() *docs.OpenApi {
	if globalSpec == nil {
		panic("spec is not initialized")
	}

	return globalSpec
}

func SetPath(paths *docs.Paths) {
	globalSpec.Paths = paths
}

func AddComponents(components *docs.Components) {
	globalSpec.Components = components
}
