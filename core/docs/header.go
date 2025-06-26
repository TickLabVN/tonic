package docs

// https://swagger.io/specification/#header-object
type HeaderObject struct {
	Description string `json:"description,omitempty"`
	Required    bool   `json:"required,omitempty"`
	Deprecated  bool   `json:"deprecated,omitempty"`
	AllowEmpty  bool   `json:"allowEmptyValue,omitempty"`

	Schema   *SchemaOrReference            `json:"schema,omitempty"`
	Style    string                        `json:"style,omitempty"`
	Explode  bool                          `json:"explode,omitempty"`
	Example  any                           `json:"example,omitempty"`
	Examples map[string]ExampleOrReference `json:"examples,omitempty"`

	Content map[string]MediaTypeObject `json:"content,omitempty" validate:"omitempty,dive"`
}

type HeaderOrReference struct {
	*HeaderObject    `json:",inline,omitempty"`
	*ReferenceObject `json:",inline,omitempty"`
}
