package docs

// Field Name	Type	Description
// schemas	Map[string, Schema Object]	An object to hold reusable Schema Objects.
// responses	Map[string, Response Object | Reference Object]	An object to hold reusable Response Objects.
// parameters	Map[string, Parameter Object | Reference Object]	An object to hold reusable Parameter Objects.
// examples	Map[string, Example Object | Reference Object]	An object to hold reusable Example Objects.
// requestBodies	Map[string, Request Body Object | Reference Object]	An object to hold reusable Request Body Objects.
// headers	Map[string, Header Object | Reference Object]	An object to hold reusable Header Objects.
// securitySchemes	Map[string, Security Scheme Object | Reference Object]	An object to hold reusable Security Scheme Objects.
// links	Map[string, Link Object | Reference Object]	An object to hold reusable Link Objects.
// callbacks	Map[string, Callback Object | Reference Object]	An object to hold reusable Callback Objects.
// pathItems	Map[string, Path Item Object | Reference Object]	An object to hold reusable Path Item Object.

type Components struct {
	Schemas         map[string]*Schema                    `json:"schemas,omitempty"`
	Responses       map[string]*ResponseOrReference       `json:"responses,omitempty"`
	Parameters      map[string]*ParameterOrReference      `json:"parameters,omitempty"`
	Examples        map[string]*ExampleOrReference        `json:"examples,omitempty"`
	RequestBodies   map[string]*RequestBodyOrReference    `json:"requestBodies,omitempty"`
	SecuritySchemes map[string]*SecuritySchemeOrReference `json:"securitySchemes,omitempty"`
	Links           map[string]*LinkOrReference           `json:"links,omitempty"`
	Callbacks       map[string]*CallbackOrReference       `json:"callbacks,omitempty"`
	PathItems       map[string]*PathItemOrReference       `json:"pathItems,omitempty"`
}
