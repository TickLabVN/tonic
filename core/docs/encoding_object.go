package docs

// https://swagger.io/specification/#encoding-object
type EncodingObject struct {
	ContentType string                       `json:"contentType,omitempty"`
	Headers     map[string]HeaderOrReference `json:"headers,omitempty"`
}
