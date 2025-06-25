package docs

// https://swagger.io/specification/#oauth-flows-object
type OAuthFlowsObject struct {
	Implicit          *OAuthFlowObject `json:"implicit,omitempty"`
	Password          *OAuthFlowObject `json:"password,omitempty"`
	ClientCredentials *OAuthFlowObject `json:"clientCredentials,omitempty"`
	AuthorizationCode *OAuthFlowObject `json:"authorizationCode,omitempty"`
}
