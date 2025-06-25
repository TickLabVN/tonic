package docs

// TODO: Add Path Templating, Examples, Schema, Example, Reference Object
// https://swagger.io/specification/#parameter-object
type ParameterObject struct {
	Name        string `json:"name,omitempty" validate:"required"`                     //TODO: verify if this is path item object
	In          string `json:"in,omitempty" validate:"oneof=query header path cookie"` //TODO: verify if this is path item object
	Description string `json:"description,omitempty"`
	Required    bool   `json:"required,omitempty" validate:"required"` //TODO: verify
	Deprecated  bool   `json:"deprecated,omitempty"`
	AllowEmpty  bool   `json:"allowEmptyValue,omitempty"`

	Style         string                        `json:"style,omitempty" validate:"oneof=matrix label form simple spaceDelimited pipeDelimited deepObject"`
	Explode       bool                          `json:"explode,omitempty"`
	AllowReserved bool                          `json:"allowReserved,omitempty"`
	Schema        *SchemaObject                 `json:"schema,omitempty"`
	Example       any                           `json:"example,omitempty"`
	Examples      map[string]ExampleOrReference `json:"examples,omitempty"`

	Content map[string]MediaTypeObject `json:"content,omitempty"`
}

type ParameterOrReference struct {
	Parameter *ParameterObject `json:",inline,omitempty"`
	Reference *ReferenceObject `json:",inline,omitempty"`
}
