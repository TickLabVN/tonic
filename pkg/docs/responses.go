package docs

// Field Name	Type	Description
// default	Response Object | Reference Object	The documentation of responses other than the ones declared for specific HTTP response codes. Use this field to cover undeclared responses.
// Field Pattern	Type	Description
// HTTP Status Code	Response Object | Reference Object	Any HTTP status code can be used as the property name, but only one property per code, to describe the expected response for that HTTP status code. This field MUST be enclosed in quotation marks (for example, “200”) for compatibility between JSON and YAML. To define a range of response codes, this field MAY contain the uppercase wildcard character X. For example, 2XX represents all response codes between [200-299]. Only the following range definitions are allowed: 1XX, 2XX, 3XX, 4XX, and 5XX. If a response is defined using an explicit code, the explicit code definition takes precedence over the range definition for that code.

// TODO: Response Object, Reference Object, HTTP Status Code
type Responses struct {
	Default        *Response           `json:"default,omitempty"`
	HttpStatusCode map[string]Response `json:"httpStatusCode,omitempty"`
}
