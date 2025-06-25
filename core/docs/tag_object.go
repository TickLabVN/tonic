package docs

// Field Name	Type	Description
// name	string	REQUIRED. The name of the tag.
// description	string	A description for the tag. CommonMark syntax MAY be used for rich text representation.
// externalDocs	External Documentation Object	Additional external documentation for this tag.
type Tag struct {
	Name         string                       `json:"name" validate:"required"`
	Description  string                       `json:"description,omitempty"`
	ExternalDocs *ExternalDocumentationObject `json:"externalDocs,omitempty"`
}
