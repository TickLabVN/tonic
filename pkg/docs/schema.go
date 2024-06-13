package docs

// Field Name	Type	Description
// discriminator	Discriminator Object	Adds support for polymorphism. The discriminator is an object name that is used to differentiate between other schemas which may satisfy the payload description. See Composition and Inheritance for more details.
// xml	XML Object	This MAY be used only on properties schemas. It has no effect on root schemas. Adds additional metadata to describe the XML representation of this property.
// externalDocs	External Documentation Object	Additional external documentation for this schema.
// example	Any	A free-form property to include an example of an instance for this schema. To represent examples that cannot be naturally represented in JSON or YAML, a string value can be used to contain the example with escaping where necessary.

// Deprecated: The example property has been deprecated in favor of the JSON Schema examples keyword. Use of example is discouraged, and later versions of this specification may remove it.

// TODO: Add Discriminator, XML, ExternalDocs
type Ref struct {
	Ref string `json:"$ref,omitempty"`
}
type Schema struct {
	*Reference `json:",inline" validate:"required_unless=type,exclude_with=Type"`
	Type string `json:"type,omitempty" validate:"required,oneof=object array string integer number boolean"`

	// For primitive types
	// string: date, date-time, password, byte, binary
	// number: float, double
	// integer: int32, int64
	*Number  `json:",inline"`
	*Integer `json:",inline"`
	*String  `json:",inline"`
	*Object  `json:",inline" validate:"required_if=type object"`
	*Array   `json:",inline" validate:"required_if=type array"`

	Description string    `json:"description,omitempty"`
	Nullable    bool      `json:"nullable,omitempty"`
	ReadOnly    bool      `json:"readOnly,omitempty"`
	WriteOnly   bool      `json:"writeOnly,omitempty"`
	Deprecated  bool      `json:"deprecated,omitempty"`
	AllOf       []*Schema `json:"allOf,omitempty"`

	Discriminator *Discriminator `json:"discriminator,omitempty"`
	XML           *XML           `json:"xml,omitempty"`
	ExternalDocs  *ExternalDocs  `json:"externalDocs,omitempty"`
}

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
	Required             []string           `json:"required,omitempty" validate:"required,unique,min=1"`
	Properties           map[string]*Schema `json:"properties,omitempty"`
	AdditionalProperties *Schema            `json:"additionalProperties,omitempty"`
}

type Array struct {
	Items   *Schema `json:"items,omitempty"`
	Minimum int     `json:"minimum,omitempty"`
	Maximum int     `json:"maximum,omitempty"`
}
