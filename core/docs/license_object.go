package docs

// Field Name	Type	Description
// name	string	REQUIRED. The license name used for the API.
// identifier	string	An SPDX license expression for the API. The identifier field is mutually exclusive of the url field.
// url	string	A URL to the license used for the API. This MUST be in the form of a URL. The url field is mutually exclusive of the identifier field.

type LicenseObject struct {
	Name       string `json:"name" validate:"required"`
	Identifier string `json:"identifier,omitempty"`
	URL        string `json:"url,omitempty" validate:"omitempty,url"`
}
