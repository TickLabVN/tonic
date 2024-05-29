package docs

// Field Name	Type	Description
// description	string	A brief description of the request body. This could contain examples of use. CommonMark syntax MAY be used for rich text representation.
// content	Map[string, Media Type Object]	REQUIRED. The content of the request body. The key is a media type or media type range and the value describes it. For requests that match multiple keys, only the most specific key is applicable. e.g. text/plain overrides text/*
// required	boolean	Determines if the request body is required in the request. Defaults to false.

type RequestBody struct {
	Description string                `json:"description,omitempty"`
	Content     map[string]*MediaType `json:"content,omitempty" validate:"required"`
	Required    bool                  `json:"required,omitempty"`
}
