package docs

// https://swagger.io/specification/#info-object
type InfoObject struct {
	Title          string         `json:"title" validate:"required"`
	Summary        string         `json:"summary,omitempty"`
	Description    string         `json:"description,omitempty"`
	TermsOfService string         `json:"termsOfService,omitempty" validate:"omitempty,url"`
	Contact        *ContactObject `json:"contact,omitempty"`
	License        *LicenseObject `json:"license,omitempty"`
	Version        string         `json:"version" validate:"required"`
}
