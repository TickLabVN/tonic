package docs

// https://swagger.io/specification/#media-type-object
type MediaTypeObject struct {
	Schema   *SchemaOrReference            `json:"schema,omitempty"`
	Example  any                           `json:"example,omitempty"`
	Examples map[string]ExampleOrReference `json:"examples,omitempty"`
	Encoding map[string]EncodingObject     `json:"encoding,omitempty"`
}
