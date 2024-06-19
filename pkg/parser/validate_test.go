package parser_test

import (
	"reflect"
	"testing"

	"github.com/TickLabVN/tonic/parser"
	"github.com/stretchr/testify/assert"
)

type User struct {
	Email string `json:"email" validate:"email"`
	Name  string `json:"name" validate:"min=3,max=20"`
	Age   int    `json:"age" validate:"min=18"`
}

func TestParseTag(t *testing.T) {
	assert := assert.New(t)
	userType := reflect.TypeOf(User{})
	fieldValidations := make(map[string]*parser.ValidateTag)
	for i := 0; i < userType.NumField(); i++ {
		field := userType.Field(i)
		tag := field.Tag.Get("validate")
		tagObj, err := parser.ParseTag(tag)
		if err != nil {
			t.Fatal(err)
		}
		fieldValidations[field.Name] = tagObj
	}
	assert.Equal(true, fieldValidations["Email"].Email)
	assert.Equal("3", fieldValidations["Name"].Min)
	assert.Equal("20", fieldValidations["Name"].Max)
	assert.Equal("18", fieldValidations["Age"].Min)
}

// See all tag in https://github.com/go-playground/validator?tab=readme-ov-file#baked-in-validations