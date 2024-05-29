package docs

// Field Name	Type	Description
// discriminator	Discriminator Object	Adds support for polymorphism. The discriminator is an object name that is used to differentiate between other schemas which may satisfy the payload description. See Composition and Inheritance for more details.
// xml	XML Object	This MAY be used only on properties schemas. It has no effect on root schemas. Adds additional metadata to describe the XML representation of this property.
// externalDocs	External Documentation Object	Additional external documentation for this schema.
// example	Any	A free-form property to include an example of an instance for this schema. To represent examples that cannot be naturally represented in JSON or YAML, a string value can be used to contain the example with escaping where necessary.

// Deprecated: The example property has been deprecated in favor of the JSON Schema examples keyword. Use of example is discouraged, and later versions of this specification may remove it.

// TODO: Add Discriminator, XML, ExternalDocs
type Schema struct {
	Discriminator interface{}  `json:"discriminator,omitempty"`
	XML           interface{}  `json:"xml,omitempty"`
	ExternalDocs  *ExternalDoc `json:"externalDocs,omitempty"`
	Example       interface{}  `json:"example,omitempty"`
}
