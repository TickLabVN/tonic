package docs

// https://swagger.io/specification/#path-item-object
type PathItemObject struct {
	Summary     string                 `json:"summary,omitempty"`
	Description string                 `json:"description,omitempty"`
	Get         *Operation             `json:"get,omitempty"`
	Put         *Operation             `json:"put,omitempty"`
	Post        *Operation             `json:"post,omitempty"`
	Delete      *Operation             `json:"delete,omitempty"`
	Options     *Operation             `json:"options,omitempty"`
	Head        *Operation             `json:"head,omitempty"`
	Patch       *Operation             `json:"patch,omitempty"`
	Trace       *Operation             `json:"trace,omitempty"`
	Servers     []ServerObject         `json:"servers,omitempty"`
	Parameters  []ParameterOrReference `json:"parameters,omitempty"`
}

type PathItemOrReference struct {
	*PathItemObject  `json:",inline,omitempty"`
	*ReferenceObject `json:",inline,omitempty"`
}
