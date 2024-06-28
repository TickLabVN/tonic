package parser

import (
	"fmt"
	"reflect"

	"github.com/TickLabVN/tonic/docs"
)

const VALIDATE = "validate"
const JSON = "json"
const QUERY = "query"
const PARAM = "param"

func ParseStruct(field reflect.Type) (*docs.Schema, error) {
	schema := docs.Schema{}
	schema.Type = "object"
	return &schema, nil
}

func ParsePrimitiveField(field reflect.Type) (*docs.Schema, error) {
	schema := docs.Schema{}
	switch field.Kind() {
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Uint, reflect.Uint16, reflect.Uint32:
		schema.Type = "integer"
		schema.Integer = &docs.Integer{Format: "int32"}
	case reflect.Int64, reflect.Uint64:
		schema.Type = "integer"
		schema.Integer = &docs.Integer{Format: "int64"}
	case reflect.Float32:
		schema.Type = "number"
		schema.Number = &docs.Number{Format: "float"}
	case reflect.Float64:
		schema.Type = "number"
		schema.Number = &docs.Number{Format: "double"}
	case reflect.String:
		schema.Type = "string"
	case reflect.Bool:
		schema.Type = "boolean"
	default:
		return nil, fmt.Errorf("unsupported type: %s", field.String())
	}
	return &schema, nil
}
