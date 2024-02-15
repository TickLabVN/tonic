package tonic_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/TickLabVN/tonic"
	"github.com/stretchr/testify/assert"
)

type Child struct {
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required"`
}

type Address struct {
	Street string `json:"street" binding:"required"`
	City   string `json:"city" binding:"required"`
}

type Person struct {
	Name    string   `json:"name" binding:"required"`
	Age     int      `json:"age" binding:"required,min=18"`
	IsAdmin bool     `json:"isAdmin"`
	Salary  float64  `json:"salary"`
	Childs  []Child  `json:"childs"`
	Address *Address `json:"address,omitempty"`
}

func TestToSwaggerType(t *testing.T) {
	dataType := reflect.TypeOf(Person{})

	for i := 0; i < dataType.NumField(); i++ {
		field := dataType.Field(i)
		switch field.Name {
		case "Name":
			assert.Equal(t, "string", tonic.ToSwaggerType(field.Type))
		case "Age":
			assert.Equal(t, "integer", tonic.ToSwaggerType(field.Type))
		case "IsAdmin":
			assert.Equal(t, "boolean", tonic.ToSwaggerType(field.Type))
		case "Salary":
			assert.Equal(t, "number", tonic.ToSwaggerType(field.Type))
		case "Childs":
			assert.Equal(t, "array", tonic.ToSwaggerType(field.Type))
		case "Address":
			assert.Equal(t, "object", tonic.ToSwaggerType(field.Type))
		}
	}
}

func TestParseBindingTagWithNumberType(t *testing.T) {
	bindingTag := "required,min=5,max=10,len=10,email,number,url"
	schema := make(map[string]interface{})
	schema["type"] = "number"
	tonic.ParseBindingTag(bindingTag, &schema)
	assert.Equal(t, 5, schema["minimum"])
	assert.Equal(t, 10, schema["maximum"])
	assert.Equal(t, true, schema["required"])
}

func TestParseBindingTagWithStringType(t *testing.T) {
	bindingTag := "required,min=5,max=10,len=10,email"
	schema := make(map[string]interface{})
	schema["type"] = "string"

	tonic.ParseBindingTag(bindingTag, &schema)

	assert.Equal(t, 5, schema["minLength"])
	assert.Equal(t, 10, schema["maxLength"])
	assert.Equal(t, true, schema["required"])
	assert.Equal(t, "email", schema["format"])
	assert.Equal(t, 10, schema["length"])
}

func TestGentonic(t *testing.T) {
	// tonic.Gentonic(Person{})
	schema := tonic.ToSwaggerSchema(reflect.TypeOf(Person{}))
	bytes, _ := json.Marshal(schema)
	t.Log(string(bytes))
}
