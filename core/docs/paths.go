package docs

import "github.com/TickLabVN/tonic/core/utils"

type Paths map[string]PathItemObject

func (p Paths) Update(path string, item PathItemObject) {
	pathItem, ok := p[path]
	if !ok {
		pathItem = PathItemObject{}
	}
	pathItem = utils.MergeStructs(pathItem, item)
	p[path] = pathItem
}
