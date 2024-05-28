package schema

type Contact struct{
	Name string `json:"name,omitempty"`
	URL string `json:"url,omitempty"`
	Email string `json:"email,omitempty"`
}
type License struct{
	Name string `json:"name" validate:"required"`
	Identifier string `json:"identifier,omitempty"`
	URL string `json:"url,omitempty" validate:"omitempty,url"`
}

type Info struct {
	Title string `json:"title" validate:"required"`
	Summary string `json:"summary,omitempty"`
	Description string `json:"description,omitempty"`
	TermsOfService string `json:"termsOfService,omitempty" validate:"omitempty,url"`
	Contact *Contact `json:"contact,omitempty"`
	License *License `json:"license,omitempty"`
	Version string   `json:"version" validate:"required"`
}
