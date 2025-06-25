package docs

// Field Pattern	Type	Description
// {expression}	Path Item Object | Reference Object	A Path Item Object, or a reference to one, used to define a callback request and expected responses. A complete example is available.

// TODO: Path Item Object, Reference Object
type CallbackObject map[string]ParameterOrReference

type CallbackOrReference struct {
	Callback  *CallbackObject  `json:",inline,omitempty"`
	Reference *ReferenceObject `json:",inline,omitempty"`
}
