package docs

// https://swagger.io/specification/#request-body-object
type RequestBody struct {
	Description string                     `json:"description,omitempty"`
	Content     map[string]MediaTypeObject `json:"content,omitempty" validate:"required"`
	Required    bool                       `json:"required,omitempty"`
}

type RequestBodyOrReference struct {
	RequestBody *RequestBody     `json:",inline,omitempty"`
	Reference   *ReferenceObject `json:",inline,omitempty"`
}
