package schema

type ServerVariableObject struct {
	Enum 	  []string `json:"enum,omitempty" validate:"omitempty,dive,required"`
	Default   string   `json:"default"`
	Description string `json:"description,omitempty"`
}

type Server struct {
	URL 	   string `json:"url" validate:"required,url"`
	Description string `json:"description,omitempty"`
	Variables   map[string]ServerVariableObject `json:"variables,omitempty"`
}