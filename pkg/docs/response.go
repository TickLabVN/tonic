package docs

// description	string	REQUIRED. A description of the response. CommonMark syntax MAY be used for rich text representation.
// headers	Map[string, Header Object | Reference Object]	Maps a header name to its definition. [RFC7230] states header names are case insensitive. If a response header is defined with the name "Content-Type", it SHALL be ignored.
// content	Map[string, Media Type Object]	A map containing descriptions of potential response payloads. The key is a media type or media type range and the value describes it. For responses that match multiple keys, only the most specific key is applicable. e.g. text/plain overrides text/*
// links	Map[string, Link Object | Reference Object]	A map of operations links that can be followed from the response. The key of the map is a short name for the link, following the naming constraints of the names for Component Objects.

// TODO: Header Object, Reference Object, Media Type Object, Link Object, Reference Object
type Response struct {
	Description string                 `json:"description,omitempty" validate:"required"`
	Headers     map[string]interface{} `json:"headers,omitempty"`
	Content     map[string]*MediaType  `json:"content,omitempty" validate:"required"`
	Links       map[string]interface{} `json:"links,omitempty"`
}
