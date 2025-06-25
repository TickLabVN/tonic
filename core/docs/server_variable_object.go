package docs

// https://swagger.io/specification/#server-variable-object
type ServerVariableObject struct {
	Enum        []string `json:"enum,omitempty" validate:"omitempty,dive,required"`
	Default     string   `json:"default" validate:"required"`
	Description string   `json:"description,omitempty"`
}