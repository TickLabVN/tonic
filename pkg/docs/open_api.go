package docs

// openapi	string	REQUIRED. This string MUST be the version number of the OpenAPI Specification that the OpenAPI document uses. The openapi field SHOULD be used by tooling to interpret the OpenAPI document. This is not related to the API info.version string.
// info	Info Object	REQUIRED. Provides metadata about the API. The metadata MAY be used by tooling as required.
// jsonSchemaDialect	string	The default value for the $schema keyword within Schema Objects contained within this OAS document. This MUST be in the form of a URI.
// servers	[Server Object]	An array of Server Objects, which provide connectivity information to a target server. If the servers property is not provided, or is an empty array, the default value would be a Server Object with a url value of /.
// paths	Paths Object	The available paths and operations for the API.
// webhooks	Map[string, Path Item Object | Reference Object] ]	The incoming webhooks that MAY be received as part of this API and that the API consumer MAY choose to implement. Closely related to the callbacks feature, this section describes requests initiated other than by an API call, for example by an out of band registration. The key name is a unique string to refer to each webhook, while the (optionally referenced) Path Item Object describes a request that may be initiated by the API provider and the expected responses. An example is available.
// components	Components Object	An element to hold various schemas for the document.
// security	[Security Requirement Object]	A declaration of which security mechanisms can be used across the API. The list of values includes alternative security requirement objects that can be used. Only one of the security requirement objects need to be satisfied to authorize a request. Individual operations can override this definition. To make security optional, an empty security requirement ({}) can be included in the array.
// tags	[Tag Object]	A list of tags used by the document with additional metadata. The order of the tags can be used to reflect on their order by the parsing tools. Not all tags that are used by the Operation Object must be declared. The tags that are not declared MAY be organized randomly or based on the tools’ logic. Each tag name in the list MUST be unique.
// externalDocs	External Documentation Object	Additional external documentation.

// TODO: Add Info, Paths, Components, Security, Tags, ExternalDocs
type OpenApi struct {
	OpenAPI           string                          `json:"openapi" validate:"required"`
	Info              Info                            `json:"info" validate:"required"`
	JsonSchemaDialect string                          `json:"jsonSchemaDialect,omitempty"`
	Servers           []*Server                       `json:"servers,omitempty"`
	Paths             *Paths                          `json:"paths"`
	Webhooks          map[string]*PathItemOrReference `json:"webhooks,omitempty"`
	Components        *Components                     `json:"components,omitempty"`
	Security          []*SecurityRequirement          `json:"security,omitempty"`
	Tags              []*Tag                          `json:"tags,omitempty"`
	ExternalDocs      *ExternalDocs                   `json:"externalDocs,omitempty"`
}
