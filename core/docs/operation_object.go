package docs

// TODO: Add ExternalDocs, Parameter Object, Request Body Object, Responses Object, Callback Object, Security Requirement Object, Server Object
// https://swagger.io/specification/#operation-object
type OperationObject struct {
	Tags         []string                       `json:"tags,omitempty"`
	Summary      string                         `json:"summary,omitempty"`
	Description  string                         `json:"description,omitempty"`
	ExternalDocs *ExternalDocumentationObject   `json:"externalDocs,omitempty"`
	OperationId  string                         `json:"operationId,omitempty"`
	Parameters   []ParameterOrReference         `json:"parameters,omitempty"`
	RequestBody  *RequestBodyOrReference        `json:"requestBody,omitempty"`
	Responses    []Response                     `json:"responses,omitempty"`
	Callbacks    map[string]CallbackOrReference `json:"callbacks,omitempty"`
	Deprecated   bool                           `json:"deprecated,omitempty"`
	Security     []SecurityRequirement          `json:"security,omitempty"`
	Servers      []ServerObject                 `json:"servers,omitempty"`
}
