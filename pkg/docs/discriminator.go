package docs

// Field Name	Type	Description
// propertyName	string	REQUIRED. The name of the property in the payload that will hold the discriminator value.
// mapping	Map[string, string]	An object to hold mappings between payload values and schema names or references.

type Discriminator struct {
	PropertyName string            `json:"propertyName" validate:"required"`
	Mapping      map[string]string `json:"mapping,omitempty"`
}
