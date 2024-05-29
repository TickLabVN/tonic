package docs

// Field Name	Type	Description
// name	string	REQUIRED. The name of the parameter. Parameter names are case sensitive.
// If in is "path", the name field MUST correspond to a template expression occurring within the path field in the Paths Object. See Path Templating for further information.
// If in is "header" and the name field is "Accept", "Content-Type" or "Authorization", the parameter definition SHALL be ignored.
// For all other cases, the name corresponds to the parameter name used by the in property.
// in	string	REQUIRED. The location of the parameter. Possible values are "query", "header", "path" or "cookie".
// description	string	A brief description of the parameter. This could contain examples of use. CommonMark syntax MAY be used for rich text representation.
// required	boolean	Determines whether this parameter is mandatory. If the parameter location is "path", this property is REQUIRED and its value MUST be true. Otherwise, the property MAY be included and its default value is false.
// deprecated	boolean	Specifies that a parameter is deprecated and SHOULD be transitioned out of usage. Default value is false.
// allowEmptyValue	boolean	Sets the ability to pass empty-valued parameters. This is valid only for query parameters and allows sending a parameter with an empty value. Default value is false. If style is used, and if behavior is n/a (cannot be serialized), the value of allowEmptyValue SHALL be ignored. Use of this property is NOT RECOMMENDED, as it is likely to be removed in a later revision.

// The rules for serialization of the parameter are specified in one of two ways. For simpler scenarios, a schema and style can describe the structure and syntax of the parameter.

// Field Name	Type	Description
// style	string	Describes how the parameter value will be serialized depending on the type of the parameter value. Default values (based on value of in): for query - form; for path - simple; for header - simple; for cookie - form.
// explode	boolean	When this is true, parameter values of type array or object generate separate parameters for each value of the array or key-value pair of the map. For other types of parameters this property has no effect. When style is form, the default value is true. For all other styles, the default value is false.
// allowReserved	boolean	Determines whether the parameter value SHOULD allow reserved characters, as defined by [RFC3986] :/?#[]@!$&'()*+,;= to be included without percent-encoding. This property only applies to parameters with an in value of query. The default value is false.
// schema	Schema Object	The schema defining the type used for the parameter.
// example	Any	Example of the parameter’s potential value. The example SHOULD match the specified schema and encoding properties if present. The example field is mutually exclusive of the examples field. Furthermore, if referencing a schema that contains an example, the example value SHALL override the example provided by the schema. To represent examples of media types that cannot naturally be represented in JSON or YAML, a string value can contain the example with escaping where necessary.
// examples	Map[ string, Example Object | Reference Object]	Examples of the parameter’s potential value. Each example SHOULD contain a value in the correct format as specified in the parameter encoding. The examples field is mutually exclusive of the example field. Furthermore, if referencing a schema that contains an example, the examples value SHALL override the example provided by the schema.

// Field Name	Type	Description
// content	Map[string, Media Type Object]	A map containing the representations for the parameter. The key is the media type and the value describes it. The map MUST only contain one entry.

// TODO: Add Path Templating, Examples, Schema, Example, Reference Object

type Parameter struct {
	Name        string `json:"name,omitempty" validate:"required"`                     //TODO: verify if this is path item object
	In          string `json:"in,omitempty" validate:"oneof=query header path cookie"` //TODO: verify if this is path item object
	Description string `json:"description,omitempty"`
	Required    bool   `json:"required,omitempty" validate:"required"` //TODO: verify
	Deprecated  bool   `json:"deprecated,omitempty"`
	AllowEmpty  bool   `json:"allowEmptyValue,omitempty"`

	Style         string              `json:"style,omitempty" validate:"oneof=matrix label form simple spaceDelimited pipeDelimited deepObject"`
	Explode       bool                `json:"explode,omitempty"`
	AllowReserved bool                `json:"allowReserved,omitempty"`
	Schema        *Schema             `json:"schema,omitempty"`
	Example       interface{}         `json:"example,omitempty"`
	Examples      map[string]*Example `json:"examples,omitempty"`

	Content map[string]*MediaType `json:"content,omitempty"`
}

// style	explode	empty	string	array	object
// matrix	false	;color	;color=blue	;color=blue,black,brown	;color=R,100,G,200,B,150
// matrix	true	;color	;color=blue	;color=blue;color=black;color=brown	;R=100;G=200;B=150
// label	false	.	.blue	.blue.black.brown	.R.100.G.200.B.150
// label	true	.	.blue	.blue.black.brown	.R=100.G=200.B=150
// form	false	color=	color=blue	color=blue,black,brown	color=R,100,G,200,B,150
// form	true	color=	color=blue	color=blue&color=black&color=brown	R=100&G=200&B=150
// simple	false	n/a	blue	blue,black,brown	R,100,G,200,B,150
// simple	true	n/a	blue	blue,black,brown	R=100,G=200,B=150
// spaceDelimited	false	n/a	n/a	blue%20black%20brown	R%20100%20G%20200%20B%20150
// pipeDelimited	false	n/a	n/a	blue|black|brown	R|100|G|200|B|150
// deepObject	true	n/a	n/a	n/a	color[R]=100&color[G]=200&color[B]=150
