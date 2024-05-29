package schema

// Field Name	Type	Description
// enum	[string]	An enumeration of string values to be used if the substitution options are from a limited set. The array MUST NOT be empty.
// default	string	REQUIRED. The default value to use for substitution, which SHALL be sent if an alternate value is not supplied. Note this behavior is different than the Schema Object’s treatment of default values, because in those cases parameter values are optional. If the enum is defined, the value MUST exist in the enum’s values.
// description	string	An optional description for the server variable. CommonMark syntax MAY be used for rich text representation.

type ServerVariableObject struct {
	Enum        []string `json:"enum,omitempty" validate:"omitempty,dive,required"`
	Default     string   `json:"default" validate:"required"`
	Description string   `json:"description,omitempty"`
}
