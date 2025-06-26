package docs

type Paths map[string]PathItemObject

func (p Paths) AddPath(path string, item PathItemObject) {
	p[path] = item
}
