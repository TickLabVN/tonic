package docs

// Field Name	Type	Description
// summary	string	Short description for the example.
// description	string	Long description for the example. CommonMark syntax MAY be used for rich text representation.
// value	Any	Embedded literal example. The value field and externalValue field are mutually exclusive. To represent examples of media types that cannot naturally represented in JSON or YAML, use a string value to contain the example, escaping where necessary.
// externalValue	string	A URI that points to the literal example. This provides the capability to reference examples that cannot easily be included in JSON or YAML documents. The value field and externalValue field are mutually exclusive. See the rules for resolving Relative References.

type Example struct {
	Summary       string `json:"summary,omitempty"`
	Description   string `json:"description,omitempty"`
	Value         any    `json:"value,omitempty"`
	ExternalValue string `json:"externalValue,omitempty"`
}

type ExampleOrReference struct {
	Example   *Example   `json:",inline,omitempty"`
	Reference *Reference `json:",inline,omitempty"`
}
