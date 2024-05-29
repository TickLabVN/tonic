package docs

// Field Name	Type	Applies To	Description
// authorizationUrl	string	oauth2 ("implicit", "authorizationCode")	REQUIRED. The authorization URL to be used for this flow. This MUST be in the form of a URL. The OAuth2 standard requires the use of TLS.
// tokenUrl	string	oauth2 ("password", "clientCredentials", "authorizationCode")	REQUIRED. The token URL to be used for this flow. This MUST be in the form of a URL. The OAuth2 standard requires the use of TLS.
// refreshUrl	string	oauth2	The URL to be used for obtaining refresh tokens. This MUST be in the form of a URL. The OAuth2 standard requires the use of TLS.
// scopes	Map[string, string]	oauth2	REQUIRED. The available scopes for the OAuth2 security scheme. A map between the scope name and a short description for it. The map MAY be empty.

type OAuthFlow struct {
	AuthorizationURL string            `json:"authorizationUrl" validate:"required,url"`
	TokenURL         string            `json:"tokenUrl" validate:"required,url"`
	RefreshURL       string            `json:"refreshUrl,omitempty" validate:"url"`
	Scopes           map[string]string `json:"scopes" validate:"required"`
}
