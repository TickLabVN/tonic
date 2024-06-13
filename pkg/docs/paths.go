package docs

type Paths map[string]*PathItem

func (p Paths) AddPath(path string, item *PathItem) {
	p[path] = item
}

func NewPaths() Paths {
	return make(Paths)
}
