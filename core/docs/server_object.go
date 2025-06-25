package docs

// https://swagger.io/specification/#server-object
type ServerObject struct {
	Url         string                    `json:"url" validate:"required,url"`
	Description string                    `json:"description,omitempty"`
	Variables   map[string]ServerVariableObject `json:"variables,omitempty"`
}
