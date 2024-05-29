package docs

// The Header Object follows the structure of the Parameter Object with the following changes:

// name MUST NOT be specified, it is given in the corresponding headers map.
// in MUST NOT be specified, it is implicitly in header.
// All traits that are affected by the location MUST be applicable to a location of header (for example, style).

// type Parameter struct {
// 	Name        string `json:"name,omitempty" validate:"required"`                     //TODO: verify if this is path item object
// 	In          string `json:"in,omitempty" validate:"oneof=query header path cookie"` //TODO: verify if this is path item object
// 	Description string `json:"description,omitempty"`
// 	Required    bool   `json:"required,omitempty" validate:"required"` //TODO: verify
// 	Deprecated  bool   `json:"deprecated,omitempty"`
// 	AllowEmpty  bool   `json:"allowEmptyValue,omitempty"`
// }

type Header struct {
	Description string `json:"description,omitempty"`
	Required    bool   `json:"required,omitempty" validate:"required"` //TODO: verify
	Deprecated  bool   `json:"deprecated,omitempty"`
	AllowEmpty  bool   `json:"allowEmptyValue,omitempty"`

	Style         string              `json:"style,omitempty" validate:"oneof=matrix label form simple spaceDelimited pipeDelimited deepObject"`
	Explode       bool                `json:"explode,omitempty"`
	AllowReserved bool                `json:"allowReserved,omitempty"`
	Schema        *Schema             `json:"schema,omitempty"`
	Example       any                 `json:"example,omitempty"`
	Examples      map[string]*Example `json:"examples,omitempty"`

	Content map[string]*MediaType `json:"content,omitempty"`
}
