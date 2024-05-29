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
	Schemas         map[string]*Schema         `json:"schemas,omitempty"`
	Responses       map[string]*Response       `json:"responses,omitempty"`
	Parameters      map[string]*Parameter      `json:"parameters,omitempty"`
	Examples        map[string]*Example        `json:"examples,omitempty"`
	RequestBodies   map[string]*RequestBody    `json:"requestBodies,omitempty"`
	SecuritySchemes map[string]*SecurityScheme `json:"securitySchemes,omitempty"`
	Links           map[string]*Link           `json:"links,omitempty"`
	Callbacks       map[string]*Callback       `json:"callbacks,omitempty"`
	PathItems       map[string]*PathItem       `json:"pathItems,omitempty"`
}
