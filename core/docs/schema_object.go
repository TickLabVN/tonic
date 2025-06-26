package docs

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/TickLabVN/tonic/core/parser"
)

type Number struct {
	Minimum float64   `json:"minimum,omitempty"`
	Maximum float64   `json:"maximum,omitempty"`
	Enum    []float64 `json:"enum,omitempty"`
	Format  string    `json:"format,omitempty"`
}

func (n *Number) Bind(v parser.ValidateFlag) {
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
	Format  string  `json:"format,omitempty"`
}

func (i *Integer) Bind(v parser.ValidateFlag) {
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
	Format  string   `json:"format,omitempty"`
}

func (s *String) Bind(v parser.ValidateFlag) {
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
	s.Format = v.GetFormat()
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

func (a *Array) Bind(v parser.ValidateFlag) {
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

func SchemaFromType(t reflect.Type) (SchemaObject, error) {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	schema := SchemaObject{}
	var err error
	schema.Type, err = toSwaggerType(t)
	if err != nil {
		return schema, err
	}

	switch schema.Type {
	case "array":
		s, err := SchemaFromType(t.Elem())
		if err != nil {
			return schema, err
		}
		schema.Items = &s
	case "object":
		schema.Object = &Object{
			Properties: make(map[string]SchemaObject),
			Required:   make([]string, 0),
		}

		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			validateTag := field.Tag.Get("validate")
			jsonAttrs := strings.Split(field.Tag.Get("json"), ",")
			if len(jsonAttrs) == 0 || jsonAttrs[0] == "-" {
				// Skip fields with no JSON tag or ignored fields
				continue
			}
			fieldName := jsonAttrs[0]

			propertyType, err := toSwaggerType(field.Type)
			if err != nil {
				return schema, err
			}
			validateOptions := parser.ParseValidateTag(validateTag)

			switch propertyType {
			case "array":
				item, err := SchemaFromType(field.Type.Elem())
				if err != nil {
					return schema, err
				}
				property := SchemaObject{
					Type: "array",
					Array: &Array{
						Items: &item,
					},
				}
				if validateOptions != nil {
					property.Array.Bind(*validateOptions)
				}
				schema.Properties[fieldName] = property
			default:
				schema.Properties[fieldName], err = SchemaFromType(field.Type)
				if err != nil {
					return schema, err
				}
			}
		}
	}
	return schema, nil
}

func toSwaggerType(t reflect.Type) (string, error) {
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
	case reflect.Map, reflect.Struct, reflect.Interface:
		return "object", nil
	case reflect.Pointer:
		return toSwaggerType(t.Elem())
	default:
		return "", fmt.Errorf("unsupported type: %s", t.Kind())
	}
}
