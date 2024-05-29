package docs

// $ref	string	Allows for a referenced definition of this path item. The referenced structure MUST be in the form of a Path Item Object. In case a Path Item Object field appears both in the defined object and the referenced object, the behavior is undefined. See the rules for resolving Relative References.
// summary	string	An optional, string summary, intended to apply to all operations in this path.
// description	string	An optional, string description, intended to apply to all operations in this path. CommonMark syntax MAY be used for rich text representation.
// get	Operation Object	A definition of a GET operation on this path.
// put	Operation Object	A definition of a PUT operation on this path.
// post	Operation Object	A definition of a POST operation on this path.
// delete	Operation Object	A definition of a DELETE operation on this path.
// options	Operation Object	A definition of a OPTIONS operation on this path.
// head	Operation Object	A definition of a HEAD operation on this path.
// patch	Operation Object	A definition of a PATCH operation on this path.
// trace	Operation Object	A definition of a TRACE operation on this path.
// servers	[Server Object]	An alternative server array to service all operations in this path.
// parameters	[Parameter Object | Reference Object]	A list of parameters that are applicable for all the operations described under this path. These parameters can be overridden at the operation level, but cannot be removed there. The list MUST NOT include duplicated parameters. A unique parameter is defined by a combination of a name and location. The list can use the Reference Object to link to parameters that are defined at the OpenAPI Object’s components/parameters.

type Path struct {
	Ref         string     `json:"$ref,omitempty"`
	Summary     string     `json:"summary,omitempty"`
	Description string     `json:"description,omitempty"`
	Get         *Operation `json:"get,omitempty"`
}
