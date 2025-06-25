package docs

// https://swagger.io/specification/#oauth-flow-object
type OAuthFlowObject struct {
	AuthorizationURL string            `json:"authorizationUrl" validate:"required,url"`
	TokenURL         string            `json:"tokenUrl" validate:"required,url"`
	RefreshURL       string            `json:"refreshUrl,omitempty" validate:"url"`
	Scopes           map[string]string `json:"scopes" validate:"required"`
}
