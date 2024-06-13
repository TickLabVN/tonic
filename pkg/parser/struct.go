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

// func ParseSchema(obj interface{}) (*docs.Schema, error) {
// 	// object array string integer number boolean
// 	dataType := reflect.TypeOf(obj)
// 	schema := docs.Schema{}

// 	switch dataType.Kind() {
// 	case reflect.Struct:
// 		for i := 0; i < dataType.NumField(); i++ {
// 			schema.Type = "object"
// 			schema.Properties = make(map[string]*docs.Schema)

// 			field := dataType.Field(i)
// 			jsonTag := field.Tag.Get(JSON)

// 			var fieldName string = ""
// 			if jsonTag != "" {
// 				fieldName = strings.Split(jsonTag, ",")[0]
// 			}
// 			if len(fieldName) == 0 {
// 				fieldName = field.Name
// 			}
// 			fieldSchema, err := ParseStructField(field)
// 			if err != nil {
// 				return nil, err
// 			}
// 			schema.Properties[fieldName] = fieldSchema
// 		}
// 	case reflect.Slice, reflect.Array:
// 		return nil, errors.ErrUnimplemented
// 	case reflect.Map:
// 		return nil, errors.ErrUnimplemented
// 	case reflect.Ptr:
// 		return nil, errors.ErrUnimplemented

// 	default:
// 		return ParsePrimitiveField(dataType)
// 	}
// 	return &schema, nil
// }

func ParseStruct(field reflect.Type) (*docs.Schema, error) {
	schema := docs.Schema{}
	schema.Type = "object"

	// schema.Properties = make(map[string]*docs.Schema)
	// for i := 0; i < field.NumField(); i++ {
	// 	field := field.Field(i)
	// 	jsonTags := strings.Split(strings.TrimSpace(field.Tag.Get(JSON)), ",")

	// 	if len(jsonTags) > 0 {
	// 		fieldName = jsonTags[0]
	// 	} else {
	// 		fieldName = field.Name
	// 	}
	// }
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
