package docs

import (
	"fmt"
	"maps"
	"reflect"
	"strconv"
	"strings"
)

type Number struct {
	Minimum float64 `json:"minimum,omitempty"`
	Maximum float64 `json:"maximum,omitempty"`
}

func (n *Number) Bind(v ValidateFlag) {
	if v.Min != "" {
		n.Minimum, _ = strconv.ParseFloat(v.Min, 64)
	}
	if v.Max != "" {
		n.Maximum, _ = strconv.ParseFloat(v.Max, 64)
	}
}

type Integer struct {
	Minimum int `json:"minimum,omitempty"`
	Maximum int `json:"maximum,omitempty"`
}

func (i *Integer) Bind(v ValidateFlag) {
	if v.Min != "" {
		i.Minimum, _ = strconv.Atoi(v.Min)
	}
	if v.Max != "" {
		i.Maximum, _ = strconv.Atoi(v.Max)
	}
}

type String struct {
	MinLength int    `json:"minLength,omitempty"`
	MaxLength int    `json:"maxLength,omitempty"`
	Pattern   string `json:"pattern,omitempty"`
}

func (s *String) Bind(v ValidateFlag) {
	if v.Min != "" {
		s.MinLength, _ = strconv.Atoi(v.Min)
	}
	if v.Max != "" {
		s.MaxLength, _ = strconv.Atoi(v.Max)
	}
	s.Pattern = v.GetPattern()
}

type Object struct {
	Required             []string                `json:"required,omitempty" validate:"required,unique,min=1"`
	Properties           map[string]SchemaObject `json:"properties,omitempty"`
	AdditionalProperties *SchemaObject           `json:"additionalProperties,omitempty"`
}

type Array struct {
	Items    *SchemaObject `json:"items,omitempty"`
	MinItems int           `json:"minItems,omitempty"`
	MaxItems int           `json:"maxItems,omitempty"`
}

func (a *Array) Bind(v ValidateFlag) {
	if v.Min != "" {
		a.MinItems, _ = strconv.Atoi(v.Min)
	}
	if v.Max != "" {
		a.MaxItems, _ = strconv.Atoi(v.Max)
	}
}

type SchemaObject struct {
	Discriminator *Discriminator               `json:"discriminator,omitempty"`
	XML           *XmlObject                   `json:"xml,omitempty"`
	ExternalDocs  *ExternalDocumentationObject `json:"externalDocs,omitempty"`

	// Deprecated: Use examples instead
	Example  any                           `json:"example,omitempty" validate:"omitempty"`
	Examples map[string]ExampleOrReference `json:"examples,omitempty"`

	Format           string `json:"format,omitempty"`
	ContentMediaType string `json:"contentMediaType,omitempty"`
	ContentEncoding  string `json:"contentEncoding,omitempty"`
	ContentSchema    string `json:"contentSchema,omitempty"`
	ReadOnly         bool   `json:"readOnly,omitempty"`
	WriteOnly        bool   `json:"writeOnly,omitempty"`

	Type     string `json:"type,omitempty" validate:"required_without=ReferenceObject AllOf AnyOf OneOf,oneof=object array string integer number boolean"`
	*Number  `json:",inline"`
	*Integer `json:",inline"`
	*String  `json:",inline"`
	*Object  `json:",inline" validate:"required_if=type object"`
	*Array   `json:",inline" validate:"required_if=type array"`
	Enum     []any `json:"enum,omitempty" validate:"omitempty,unique"` // Enum values for string, number, integer types

	AllOf []SchemaOrReference `json:"allOf,omitempty"`
	AnyOf []SchemaOrReference `json:"anyOf,omitempty"`
	OneOf []SchemaOrReference `json:"oneOf,omitempty"`

	Description string `json:"description,omitempty"`
	Nullable    bool   `json:"nullable,omitempty"`
	Deprecated  bool   `json:"deprecated,omitempty"`

	*ReferenceObject `json:",inline" validate:"required_without=type,exclude_with=Type"`
}

type SchemaOrReference struct {
	*SchemaObject    `json:",inline,omitempty"`
	*ReferenceObject `json:",inline,omitempty"`
}

// Gin framework use "binding" tag, for example: `binding:"required,min=1,max=10"`
func SchemaFromType(t reflect.Type, parsingKey string, validateKey string, flag *ValidateFlag) (SchemaObject, error) {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	schema := SchemaObject{}
	var err error
	schema.Type, err = toSwaggerType(t)
	if err != nil {
		return schema, err
	}
	if schema.Type == "map" {
		schema.Type = "object"
		additionalProp, err := SchemaFromType(t.Elem(), parsingKey, validateKey, flag)
		if err != nil {
			return schema, err
		}
		schema.Object = &Object{
			AdditionalProperties: &additionalProp,
		}
		return schema, nil
	}

	switch schema.Type {
	case "datetime":
		schema.Type = "string"
		schema.Format = "date-time"
	case "integer":
		schema.Integer = &Integer{}
		if flag != nil {
			schema.Integer.Bind(*flag)
			if len(flag.Eq) > 0 || len(flag.OneOf) > 0 {
				schema.Enum = make([]any, 0, len(flag.Eq)+len(flag.OneOf))
				for _, val := range flag.Eq {
					valFloat, _ := strconv.Atoi(val)
					schema.Enum = append(schema.Enum, valFloat)
				}
				for _, val := range flag.OneOf {
					valFloat, _ := strconv.Atoi(val)
					schema.Enum = append(schema.Enum, valFloat)
				}
			}
		}
		schema.Format = REFLECT_TYPE_MAP[t.Kind()]
	case "number":
		schema.Number = &Number{}
		if flag != nil {
			schema.Number.Bind(*flag)
			if len(flag.Eq) > 0 || len(flag.OneOf) > 0 {
				schema.Enum = make([]any, 0, len(flag.Eq)+len(flag.OneOf))
				for _, val := range flag.Eq {
					valFloat, _ := strconv.ParseFloat(val, 64)
					schema.Enum = append(schema.Enum, valFloat)
				}
				for _, val := range flag.OneOf {
					valFloat, _ := strconv.ParseFloat(val, 64)
					schema.Enum = append(schema.Enum, valFloat)
				}
			}
		}
		schema.Format = REFLECT_TYPE_MAP[t.Kind()]
	case "string":
		schema.String = &String{}
		if flag != nil {
			schema.String.Bind(*flag)
			schema.Format = flag.GetFormat()
			if len(flag.Eq) > 0 || len(flag.EqIgnoreCase) > 0 || len(flag.OneOf) > 0 {
				enums := make([]string, 0, len(flag.Eq)+len(flag.EqIgnoreCase)+len(flag.OneOf))
				enums = append(enums, flag.Eq...)
				enums = append(enums, flag.EqIgnoreCase...)
				enums = append(enums, flag.OneOf...)
				a := make([]any, len(enums))
				for i, v := range enums {
					a[i] = v
				}
				schema.Enum = a
			}
		}
	case "boolean":
		schema.Enum = make([]any, 0, 2)
		if flag != nil {
			if len(flag.Eq) > 0 {
				for _, v := range flag.Eq {
					b, err := strconv.ParseBool(v)
					if err != nil {
						return schema, fmt.Errorf("invalid boolean value: %s", v)
					}
					schema.Enum = append(schema.Enum, b)
				}
			} else {
				schema.Enum = []any{true, false}
			}
		}
	case "array":
		s, err := SchemaFromType(t.Elem(), parsingKey, validateKey, nil)
		if err != nil {
			return schema, err
		}
		schema.Array = &Array{
			Items: &s,
		}
		if flag != nil {
			schema.Array.Bind(*flag)
		}
	case "object":
		schema.Object = &Object{
			Properties: make(map[string]SchemaObject),
			Required:   make([]string, 0),
		}

		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			// For embedded structs, we need to handle them differently
			if field.Anonymous {
				embeddedSchema, err := SchemaFromType(field.Type, parsingKey, validateKey, nil)
				if err != nil {
					return schema, fmt.Errorf("create schema from type %s: %w", field.Type.String(), err)
				}
				maps.Copy(schema.Properties, embeddedSchema.Properties)
				schema.Required = append(schema.Required, embeddedSchema.Required...)
				continue
			}

			validateTag := field.Tag.Get(validateKey)
			parsingAttr := strings.Split(field.Tag.Get(parsingKey), ",")
			var fieldName string
			if len(parsingAttr) == 0 || parsingAttr[0] == "" {
				continue // Skip if no parsing attributes are provided
			}
			fieldName = parsingAttr[0]
			validateOptions, err := ParseValidateTag(validateTag)
			if err != nil {
				return schema, err
			}

			schema.Properties[fieldName], err = SchemaFromType(field.Type, parsingKey, validateKey, validateOptions)
			if err != nil {
				return schema, err
			}

			if validateOptions != nil && validateOptions.Required {
				schema.Required = append(schema.Required, fieldName)
			}
		}
	}
	return schema, nil
}

var REFLECT_TYPE_MAP = map[reflect.Kind]string{
	reflect.Int:     "int32",
	reflect.Int8:    "int8",
	reflect.Int16:   "int16",
	reflect.Int32:   "int32",
	reflect.Int64:   "int64",
	reflect.Uint:    "uint",
	reflect.Uint8:   "uint8",
	reflect.Uint16:  "uint16",
	reflect.Uint32:  "uint32",
	reflect.Uint64:  "uint64",
	reflect.Float32: "float32",
	reflect.Float64: "float64",
}

func toSwaggerType(t reflect.Type) (string, error) {
	if t.PkgPath() == "time" && t.Name() == "Time" {
		return "datetime", nil // Time is represented as datetime in OpenAPI
	}
	switch t.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return "integer", nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return "integer", nil
	case reflect.Float32, reflect.Float64:
		return "number", nil
	case reflect.String:
		return "string", nil
	case reflect.Bool:
		return "boolean", nil
	case reflect.Slice, reflect.Array:
		return "array", nil
	case reflect.Map:
		return "map", nil
	case reflect.Struct, reflect.Interface:
		return "object", nil
	case reflect.Pointer:
		return toSwaggerType(t.Elem())
	default:
		return "", fmt.Errorf("unsupported type: %s", t.Kind())
	}
}
