package docs

// Field Pattern	Type	Description
// {expression}	Path Item Object | Reference Object	A Path Item Object, or a reference to one, used to define a callback request and expected responses. A complete example is available.

// TODO: Path Item Object, Reference Object
type Callback struct {
	Expression map[string]*PathItem `json:"expression,omitempty"` // TODO: verify the Expression
}
