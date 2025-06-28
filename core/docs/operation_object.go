package docs

import (
	"slices"
)

// https://swagger.io/specification/#operation-object
type OperationObject struct {
	Tags         []string                       `json:"tags,omitempty"`
	Summary      string                         `json:"summary,omitempty"`
	Description  string                         `json:"description,omitempty"`
	ExternalDocs *ExternalDocumentationObject   `json:"externalDocs,omitempty"`
	OperationId  string                         `json:"operationId,omitempty"`
	Parameters   []ParameterOrReference         `json:"parameters,omitempty"`
	RequestBody  *RequestBodyOrReference        `json:"requestBody,omitempty"`
	Responses    map[string]ResponseOrReference `json:"responses,omitempty"`
	Callbacks    map[string]CallbackOrReference `json:"callbacks,omitempty"`
	Deprecated   bool                           `json:"deprecated,omitempty"`
	Security     []SecurityRequirement          `json:"security,omitempty"`
	Servers      []ServerObject                 `json:"servers,omitempty"`
}

func (o *OperationObject) AddParameter(in string, objectSchema *SchemaObject, schemaPath string) {
	if o.Parameters == nil {
		o.Parameters = make([]ParameterOrReference, 0)
	}

	for propName, prop := range objectSchema.Properties {
		o.Parameters = append(o.Parameters, ParameterOrReference{
			ParameterObject: &ParameterObject{
				Name:        propName,
				In:          in,
				Description: prop.Description,
				Schema: &SchemaOrReference{
					ReferenceObject: &ReferenceObject{
						Ref: schemaPath + "/properties/" + propName,
					},
				},
				Required: slices.Contains(objectSchema.Required, propName) || in == "path",
			},
		})
	}
}
