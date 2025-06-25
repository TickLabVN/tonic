package docs

// https://swagger.io/specification/#reference-object
type ReferenceObject struct {
	Ref         string `json:"$ref" validate:"required"`
	Summary     string `json:"summary,omitempty"`
	Description string `json:"description,omitempty"`
}
