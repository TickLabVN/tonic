package docs

// https://swagger.io/specification/#security-scheme-object
type SecuritySchemeObject struct {
	Type             string            `json:"type" validate:"required,oneof=apiKey http mutualTLS oauth2 openIdConnect"`
	Description      string            `json:"description,omitempty"`
	Name             string            `json:"name,omitempty" validate:"required_if=Type apiKey"`
	In               string            `json:"in,omitempty" validate:"required_if=Type apiKey,oneof=query header cookie"`
	Scheme           string            `json:"scheme,omitempty" validate:"required_if=Type http"`
	BearerFormat     string            `json:"bearerFormat,omitempty"`
	Flows            *OAuthFlowsObject `json:"flows,omitempty" validate:"required_if=Type oauth2"`
	OpenIDConnectURL string            `json:"openIdConnectUrl,omitempty" validate:"required_if=Type openIdConnect,url"`
}

type SecuritySchemeOrReference struct {
	*SecuritySchemeObject `json:",inline,omitempty"`
	*ReferenceObject      `json:",inline,omitempty"`
}

type SecurityRequirement map[string][]string
