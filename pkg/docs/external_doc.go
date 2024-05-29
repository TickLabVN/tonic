package docs

// Field Name	Type	Description
// description	string	A description of the target documentation. CommonMark syntax MAY be used for rich text representation.
// url	string	REQUIRED. The URL for the target documentation. This MUST be in the form of a URL.

type ExternalDoc struct {
	Description string `json:"description,omitempty"`
	URL         string `json:"url,omitempty" validate:"required,url"`
}
