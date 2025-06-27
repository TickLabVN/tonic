package docs

import (
	"fmt"
	"maps"
	"reflect"
	"strconv"
	"strings"
)

type Number struct {
	Minimum float64   `json:"minimum,omitempty"`
	Maximum float64   `json:"maximum,omitempty"`
	Enum    []float64 `json:"enum,omitempty"`
}

func (n *Number) Bind(v ValidateFlag) {
	if v.Min != "" {
		n.Minimum, _ = strconv.ParseFloat(v.Min, 64)
	}
	if v.Max != "" {
		n.Maximum, _ = strconv.ParseFloat(v.Max, 64)
	}
	if v.OneOf != nil {
		n.Enum = make([]float64, len(v.OneOf))
		for i, val := range v.OneOf {
			n.Enum[i], _ = strconv.ParseFloat(val, 64)
		}
	}
}

type Integer struct {
	Minimum int     `json:"minimum,omitempty"`
	Maximum int     `json:"maximum,omitempty"`
	Enum    []int64 `json:"enum,omitempty"`
}

func (i *Integer) Bind(v ValidateFlag) {
	if v.Min != "" {
		i.Minimum, _ = strconv.Atoi(v.Min)
	}
	if v.Max != "" {
		i.Maximum, _ = strconv.Atoi(v.Max)
	}
	if v.OneOf != nil {
		i.Enum = make([]int64, len(v.OneOf))
		for j, val := range v.OneOf {
			i.Enum[j], _ = strconv.ParseInt(val, 10, 64)
		}
	}
}

type String struct {
	Minimum int      `json:"minimum,omitempty"`
	Maximum int      `json:"maximum,omitempty"`
	Pattern string   `json:"pattern,omitempty"`
	Enum    []string `json:"enum,omitempty"`
}

func (s *String) Bind(v ValidateFlag) {
	if v.Min != "" {
		s.Minimum, _ = strconv.Atoi(v.Min)
	}
	if v.Max != "" {
		s.Maximum, _ = strconv.Atoi(v.Max)
	}
	s.Pattern = v.GetPattern()
	if v.OneOf != nil {
		s.Enum = v.OneOf
	}
}

type Object struct {
	Required             []string                `json:"required,omitempty" validate:"required,unique,min=1"`
	Properties           map[string]SchemaObject `json:"properties,omitempty"`
	AdditionalProperties *SchemaObject           `json:"additionalProperties,omitempty"`
}

type Array struct {
	Items   *SchemaObject `json:"items,omitempty"`
	Minimum int           `json:"minimum,omitempty"`
	Maximum int           `json:"maximum,omitempty"`
}

func (a *Array) Bind(v ValidateFlag) {
	if v.Min != "" {
		a.Minimum, _ = strconv.Atoi(v.Min)
	}
	if v.Max != "" {
		a.Maximum, _ = strconv.Atoi(v.Max)
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
func SchemaFromType(t reflect.Type, bindingKey string, flag *ValidateFlag) (SchemaObject, error) {
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
		additionalProp, err := SchemaFromType(t.Elem(), bindingKey, flag)
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
		}
		schema.Format = REFLECT_TYPE_MAP[t.Kind()]
	case "number":
		schema.Number = &Number{}
		if flag != nil {
			schema.Number.Bind(*flag)
		}
		schema.Format = REFLECT_TYPE_MAP[t.Kind()]
	case "string":
		schema.String = &String{}
		if flag != nil {
			schema.String.Bind(*flag)
			schema.Format = flag.GetFormat()
		}
	case "boolean":
		// No additional properties for boolean
	case "array":
		s, err := SchemaFromType(t.Elem(), bindingKey, nil)
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
				embeddedSchema, err := SchemaFromType(field.Type, bindingKey, nil)
				if err != nil {
					return schema, fmt.Errorf("create schema from type %s: %w", field.Type.String(), err)
				}
				maps.Copy(schema.Properties, embeddedSchema.Properties)
				schema.Required = append(schema.Required, embeddedSchema.Required...)
				continue
			}

			validateTag := field.Tag.Get(bindingKey)
			jsonAttrs := strings.Split(field.Tag.Get("json"), ",")
			var fieldName string
			if len(jsonAttrs) > 0 && jsonAttrs[0] != "" {
				fieldName = jsonAttrs[0]
			} else {
				fieldName = field.Name
			}
			validateOptions, err := ParseValidateTag(validateTag)
			if err != nil {
				return schema, fmt.Errorf("parse validate tag for field %s: %w", fieldName, err)
			}

			schema.Properties[fieldName], err = SchemaFromType(field.Type, bindingKey, validateOptions)
			if err != nil {
				return schema, fmt.Errorf("create schema from type %s: %w", field.Type.String(), err)
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
