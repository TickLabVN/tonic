package docs

// https://swagger.io/specification/#contact-object
type ContactObject struct {
	Name  string `json:"name,omitempty"`
	URL   string `json:"url,omitempty" validate:"url"`
	Email string `json:"email,omitempty" validate:"email"`
}
