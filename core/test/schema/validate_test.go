package schema_test

import (
	"reflect"
	"testing"

	"github.com/TickLabVN/tonic/core/validator"
	"github.com/stretchr/testify/assert"
)

type Education struct {
	School string `json:"school" validate:"required"`
	Degree string `json:"degree" validate:"required"`
}

type User struct {
	Email      string      `json:"email" validate:"email,required"`
	Name       string      `json:"name" validate:"min=3,max=20"`
	Age        int         `json:"age" validate:"min=18"`
	Educations []Education `json:"educations" validate:"dive,required"`
}

func TestParseTag(t *testing.T) {
	assert := assert.New(t)

	user := User{
		Email: "np@example.com",
		Name:  "Vinh",
		Age:   20,
	}
	userType := reflect.TypeOf(user)
	t.Log("User type:", userType.Name(), userType.PkgPath())

	fieldValidations := make(map[string]validator.ValidateFlag)
	for i := 0; i < userType.NumField(); i++ {
		field := userType.Field(i)
		tag := field.Tag.Get("validate")
		tagObj, err := validator.ParseValidateTag(tag)
		if err != nil {
			t.Fatal(err)
		}
		fieldValidations[field.Name] = *tagObj
	}
	assert.Equal(true, fieldValidations["Email"].Email)
	assert.Equal(true, fieldValidations["Email"].Required)
	assert.Equal("3", fieldValidations["Name"].Min)
	assert.Equal("20", fieldValidations["Name"].Max)
	assert.Equal("18", fieldValidations["Age"].Min)
}
