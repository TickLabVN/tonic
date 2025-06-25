package docs

// $ref	string	REQUIRED. The reference identifier. This MUST be in the form of a URI.
// summary	string	A short summary which by default SHOULD override that of the referenced component. If the referenced object-type does not allow a summary field, then this field has no effect.
// description	string	A description which by default SHOULD override that of the referenced component. CommonMark syntax MAY be used for rich text representation. If the referenced object-type does not allow a description field, then this field has no effect.

type Reference struct {
	Ref         string `json:"$ref" validate:"required"`
	Summary     string `json:"summary,omitempty"`
	Description string `json:"description,omitempty"`
}
