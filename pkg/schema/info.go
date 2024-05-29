package schema

// Field Name	Type	Description
// title	string	REQUIRED. The title of the API.
// summary	string	A short summary of the API.
// description	string	A description of the API. CommonMark syntax MAY be used for rich text representation.
// termsOfService	string	A URL to the Terms of Service for the API. This MUST be in the form of a URL.
// contact	Contact Object	The contact information for the exposed API.
// license	License Object	The license information for the exposed API.
// version	string	REQUIRED. The version of the OpenAPI document (which is distinct from the OpenAPI Specification version or the API implementation version).

type Info struct {
	Title          string   `json:"title" validate:"required"`
	Summary        string   `json:"summary,omitempty"`
	Description    string   `json:"description,omitempty"`
	TermsOfService string   `json:"termsOfService,omitempty" validate:"omitempty,url"`
	Contact        *Contact `json:"contact,omitempty"`
	License        *License `json:"license,omitempty"`
	Version        string   `json:"version" validate:"required"`
}
