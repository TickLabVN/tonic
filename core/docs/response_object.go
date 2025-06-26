package docs

// https://swagger.io/specification/#response-object
type ResponseObject struct {
	Description string                       `json:"description,omitempty" validate:"required"`
	Headers     map[string]HeaderOrReference `json:"headers,omitempty"`
	Content     map[string]MediaTypeObject   `json:"content,omitempty" validate:"required"`
	Links       map[string]LinkOrReference   `json:"links,omitempty"`
}

type ResponseOrReference struct {
	*ResponseObject  `json:",inline,omitempty"`
	*ReferenceObject `json:",inline,omitempty"`
}
