package docs

// Field Name	Type	Description
// name	string	REQUIRED. The name of the parameter. Parameter names are case sensitive.
// If in is "path", the name field MUST correspond to a template expression occurring within the path field in the Paths Object. See Path Templating for further information.
// If in is "header" and the name field is "Accept", "Content-Type" or "Authorization", the parameter definition SHALL be ignored.
// For all other cases, the name corresponds to the parameter name used by the in property.
// in	string	REQUIRED. The location of the parameter. Possible values are "query", "header", "path" or "cookie".
// description	string	A brief description of the parameter. This could contain examples of use. CommonMark syntax MAY be used for rich text representation.
// required	boolean	Determines whether this parameter is mandatory. If the parameter location is "path", this property is REQUIRED and its value MUST be true. Otherwise, the property MAY be included and its default value is false.
// deprecated	boolean	Specifies that a parameter is deprecated and SHOULD be transitioned out of usage. Default value is false.
// allowEmptyValue	boolean	Sets the ability to pass empty-valued parameters. This is valid only for query parameters and allows sending a parameter with an empty value. Default value is false. If style is used, and if behavior is n/a (cannot be serialized), the value of allowEmptyValue SHALL be ignored. Use of this property is NOT RECOMMENDED, as it is likely to be removed in a later revision.

// TODO: Add Path Templating
type Param struct {
	Name        string `json:"name,omitempty" validate:"required"`                     //TODO: verify if this is path item object
	In          string `json:"in,omitempty" validate:"oneof=query header path cookie"` //TODO: verify if this is path item object
	Description string `json:"description,omitempty"`
	Required    bool   `json:"required,omitempty" validate:"required"` //TODO: verify
	Deprecated  bool   `json:"deprecated,omitempty"`
	AllowEmpty  bool   `json:"allowEmptyValue,omitempty"`
}
