package docs

// Field Name	Type	Description
// url	string	REQUIRED. A URL to the target host. This URL supports Server Variables and MAY be relative, to indicate that the host location is relative to the location where the OpenAPI document is being served. Variable substitutions will be made when a variable is named in {brackets}.
// description	string	An optional string describing the host designated by the URL. CommonMark syntax MAY be used for rich text representation.
// variables	Map[string, Server Variable Object]	A map between a variable name and its value. The value is used for substitution in the server’s URL template.

type Server struct {
	URL         string                          `json:"url" validate:"required,url"`
	Description string                          `json:"description,omitempty"`
	Variables   map[string]ServerVariableObject `json:"variables,omitempty"`
}
