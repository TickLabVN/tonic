package docs

import (
	"fmt"
	"reflect"

	"github.com/TickLabVN/tonic/core/utils"
)

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

type ComponentsObject struct {
	Schemas         map[string]SchemaOrReference         `json:"schemas,omitempty"`
	Responses       map[string]ResponseOrReference       `json:"responses,omitempty"`
	Parameters      map[string]ParameterOrReference      `json:"parameters,omitempty"`
	Examples        map[string]ExampleOrReference        `json:"examples,omitempty"`
	RequestBodies   map[string]RequestBodyOrReference    `json:"requestBodies,omitempty"`
	SecuritySchemes map[string]SecuritySchemeOrReference `json:"securitySchemes,omitempty"`
	Links           map[string]LinkOrReference           `json:"links,omitempty"`
	Callbacks       map[string]CallbackOrReference       `json:"callbacks,omitempty"`
	PathItems       map[string]PathItemOrReference       `json:"pathItems,omitempty"`
}

func (c *ComponentsObject) AddSchema(t reflect.Type, bindingKey string) error {
	if c.Schemas == nil {
		c.Schemas = make(map[string]SchemaOrReference)
	}

	schemaName := utils.GetSchemaName(t)
	if _, exists := c.Schemas[schemaName]; exists {
		// Schema already exists, no need to add it again
		return nil
	}

	schema, err := SchemaFromType(t, bindingKey, nil)
	if err != nil {
		return fmt.Errorf("create schema from type: %w", err)
	}
	c.Schemas[schemaName] = SchemaOrReference{SchemaObject: &schema}
	return nil
}

func (c *ComponentsObject) AddResponse(name string, response *ResponseOrReference) {
	if c.Responses == nil {
		c.Responses = make(map[string]ResponseOrReference)
	}
	c.Responses[name] = *response
}

func (c *ComponentsObject) AddParameter(name string, parameter *ParameterOrReference) {
	if c.Parameters == nil {
		c.Parameters = make(map[string]ParameterOrReference)
	}
	c.Parameters[name] = *parameter
}

func (c *ComponentsObject) AddExample(name string, example *ExampleOrReference) {
	if c.Examples == nil {
		c.Examples = make(map[string]ExampleOrReference)
	}
	c.Examples[name] = *example
}

func (c *ComponentsObject) AddRequestBody(name string, requestBody *RequestBodyOrReference) {
	if c.RequestBodies == nil {
		c.RequestBodies = make(map[string]RequestBodyOrReference)
	}
	c.RequestBodies[name] = *requestBody
}

func (c *ComponentsObject) AddSecurityScheme(name string, securityScheme *SecuritySchemeOrReference) {
	if c.SecuritySchemes == nil {
		c.SecuritySchemes = make(map[string]SecuritySchemeOrReference)
	}
	c.SecuritySchemes[name] = *securityScheme
}

func (c *ComponentsObject) AddLink(name string, link *LinkOrReference) {
	if c.Links == nil {
		c.Links = make(map[string]LinkOrReference)
	}
	c.Links[name] = *link
}

func (c *ComponentsObject) AddCallback(name string, callback *CallbackOrReference) {
	if c.Callbacks == nil {
		c.Callbacks = make(map[string]CallbackOrReference)
	}
	c.Callbacks[name] = *callback
}

func (c *ComponentsObject) AddPathItem(name string, pathItem *PathItemOrReference) {
	if c.PathItems == nil {
		c.PathItems = make(map[string]PathItemOrReference)
	}
	c.PathItems[name] = *pathItem
}
