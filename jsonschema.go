package tonic

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func ToSwaggerSchema(t reflect.Type) map[string]interface{} {
	kind := t.Kind()
	schema := make(map[string]interface{})

	if kind == reflect.String {
		schema["type"] = "string"
	} else if kind == reflect.Int || kind == reflect.Int8 || kind == reflect.Int16 || kind == reflect.Int32 || kind == reflect.Int64 {
		schema["type"] = "integer"
	} else if kind == reflect.Float32 || kind == reflect.Float64 {
		schema["type"] = "number"
	} else if kind == reflect.Bool {
		schema["type"] = "boolean"
	} else if kind == reflect.Ptr {
		schema = ToSwaggerSchema(t.Elem())
	} else if kind == reflect.Array || kind == reflect.Slice {
		schema["type"] = "array"
		schema["items"] = ToSwaggerSchema(t.Elem())
	} else if kind == reflect.Struct {
		schema["type"] = "object"
		properties := make(map[string]interface{})
		schema["properties"] = properties
		requiredFields := []string{}

		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			jsonTag := field.Tag.Get("json")
			jsonAttrs := strings.Split(jsonTag, ",")
			fieldName := jsonAttrs[0]

			if fieldName == "" {
				continue
			}

			fieldKind := field.Type.Kind()
			if fieldKind == reflect.Pointer {
				properties[fieldName] = ToSwaggerSchema(field.Type.Elem())
				continue
			} else if fieldKind == reflect.Struct {
				properties[fieldName] = ToSwaggerSchema(field.Type)
				continue
			}

			fieldSchema := make(map[string]interface{})
			fieldSchema["type"] = ToSwaggerType(field.Type)

			if fieldKind == reflect.Array || fieldKind == reflect.Slice {
				fieldSchema["items"] = ToSwaggerSchema(field.Type.Elem())
			} else {
				bindingTag := field.Tag.Get("binding")

				if len(jsonAttrs) > 1 && jsonAttrs[1] == "omitempty" {
					fieldSchema["required"] = false
				}
				isRequired := ParseBindingTag(bindingTag, &fieldSchema)
				if isRequired {
					requiredFields = append(requiredFields, fieldName)
				}
			}

			properties[fieldName] = fieldSchema
		}
		schema["required"] = requiredFields
	}

	return schema
}

func ParseBindingTag(bindingTag string, fieldSchema *map[string]interface{}) (isRequired bool) {
	tags := strings.Split(bindingTag, ",")
	fieldType, _ := (*fieldSchema)["type"].(string)
	isRequired = false
	for _, tag := range tags {
		if tag == "required" {
			isRequired = true
		} else if strings.HasPrefix(tag, "max=") {
			maxValue, err := strconv.Atoi(strings.TrimPrefix(tag, "max="))
			if err != nil {
				panic(err)
			}
			if fieldType == "integer" || fieldType == "number" {
				(*fieldSchema)["maximum"] = maxValue
			} else if fieldType == "string" {
				(*fieldSchema)["maxLength"] = maxValue
			}
		} else if strings.HasPrefix(tag, "min=") {
			minValue, err := strconv.Atoi(strings.TrimPrefix(tag, "min="))
			if err != nil {
				panic(err)
			}
			if fieldType == "integer" || fieldType == "number" {
				(*fieldSchema)["minimum"] = minValue
			} else if fieldType == "string" {
				(*fieldSchema)["minLength"] = minValue
			}
		} else if fieldType == "string" {
			if strings.HasPrefix(tag, "len=") {
				length, err := strconv.Atoi(strings.TrimPrefix(tag, "len="))
				if err != nil {
					panic(err)
				}
				(*fieldSchema)["length"] = length
			} else {
				(*fieldSchema)["format"] = tag
			}
		}
	}

	return isRequired
}

func ToSwaggerType(fieldType reflect.Type) string {
	kind := fieldType.Kind()
	switch kind {
	case reflect.String:
		return "string"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return "integer"
	case reflect.Float32, reflect.Float64:
		return "number"
	case reflect.Bool:
		return "boolean"
	case reflect.Struct:
		return "object"
	case reflect.Array, reflect.Slice:
		return "array"
	case reflect.Pointer:
		return ToSwaggerType(fieldType.Elem())
	default:
		panic(fmt.Sprintf("Unsupported type: %s", kind))
	}
}
