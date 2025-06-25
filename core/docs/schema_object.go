package docs

type Number struct {
	Minimum float64   `json:"minimum,omitempty"`
	Maximum float64   `json:"maximum,omitempty"`
	Enum    []float64 `json:"enum,omitempty"`
	Format  string    `json:"format,omitempty"`
}

type Integer struct {
	Minimum int     `json:"minimum,omitempty"`
	Maximum int     `json:"maximum,omitempty"`
	Enum    []int64 `json:"enum,omitempty"`
	Format  string  `json:"format,omitempty"`
}

type String struct {
	Minimum int      `json:"minimum,omitempty"`
	Maximum int      `json:"maximum,omitempty"`
	Pattern string   `json:"pattern,omitempty"`
	Enum    []string `json:"enum,omitempty"`
	Format  string   `json:"format,omitempty"`
}

type Object struct {
	Required             []string                `json:"required,omitempty" validate:"required,unique,min=1"`
	Properties           map[string]SchemaObject `json:"properties,omitempty"`
	AdditionalProperties *SchemaObject           `json:"additionalProperties,omitempty"`
}

type Array struct {
	Items   *SchemaObject `json:"items,omitempty"`
	Minimum int           `json:"minimum,omitempty"`
	Maximum int           `json:"maximum,omitempty"`
}

type SchemaObject struct {
	Discriminator *Discriminator               `json:"discriminator,omitempty"`
	XML           *XmlObject                   `json:"xml,omitempty"`
	ExternalDocs  *ExternalDocumentationObject `json:"externalDocs,omitempty"`

	// Deprecated: Use examples instead
	Example  any                           `json:"example,omitempty" validate:"omitempty"`
	Examples map[string]ExampleOrReference `json:"examples,omitempty"`

	Format           string `json:"format,omitempty"`
	ContentMediaType string `json:"contentMediaType,omitempty"`
	ContentEncoding  string `json:"contentEncoding,omitempty"`
	ContentSchema    string `json:"contentSchema,omitempty"`
	ReadOnly         bool   `json:"readOnly,omitempty"`
	WriteOnly        bool   `json:"writeOnly,omitempty"`

	Type     string `json:"type,omitempty" validate:"required_without=ReferenceObject AllOf AnyOf OneOf,oneof=object array string integer number boolean"`
	*Number  `json:",inline"`
	*Integer `json:",inline"`
	*String  `json:",inline"`
	*Object  `json:",inline" validate:"required_if=type object"`
	*Array   `json:",inline" validate:"required_if=type array"`

	AllOf []SchemaOrReference `json:"allOf,omitempty"`
	AnyOf []SchemaOrReference `json:"anyOf,omitempty"`
	OneOf []SchemaOrReference `json:"oneOf,omitempty"`

	Description string `json:"description,omitempty"`
	Nullable    bool   `json:"nullable,omitempty"`
	Deprecated  bool   `json:"deprecated,omitempty"`

	*ReferenceObject `json:",inline" validate:"required_without=type,exclude_with=Type"`
}

type SchemaOrReference struct {
	Schema    *SchemaObject    `json:",inline,omitempty"`
	Reference *ReferenceObject `json:",inline,omitempty"`
}
