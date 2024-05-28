package schema

type ExternalDocumentationObject struct {
	Url         string `json:"url" validate:"required,url"`
	Description string `json:"description,omitempty"`
}

type Tag struct {
	Name         string                       `json:"name" validate:"required"`
	Description  string                       `json:"description,omitempty"`
	ExternalDocs *ExternalDocumentationObject `json:"externalDocs,omitempty"`
}
