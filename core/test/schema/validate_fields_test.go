package schema_test

import (
	"testing"

	"github.com/TickLabVN/tonic/core/utils"
	"github.com/stretchr/testify/assert"
)

// eqcsfield	Field Equals Another Field (relative)
// eqfield	Field Equals Another Field
// fieldcontains	Check the indicated characters are present in the Field
// fieldexcludes	Check the indicated characters are not present in the field
// gtcsfield	Field Greater Than Another Relative Field
// gtecsfield	Field Greater Than or Equal To Another Relative Field
// gtefield	Field Greater Than or Equal To Another Field
// gtfield	Field Greater Than Another Field
// ltcsfield	Less Than Another Relative Field
// ltecsfield	Less Than or Equal To Another Relative Field
// ltefield	Less Than or Equal To Another Field
// ltfield	Less Than Another Field
// necsfield	Field Does Not Equal Another Field (relative)
// nefield	Field Does Not Equal Another Field

func TestValidate_Field(t *testing.T) {
	type Test struct {
		EQCSField     string `validate:"eqcsfield=Field"`
		EQField       string `validate:"eqfield=Field"`
		FieldContains string `validate:"fieldcontains=Field"`
		FieldExcludes string `validate:"fieldexcludes=Field"`
		GTCSField     string `validate:"gtcsfield=Field"`
		GTECSField    string `validate:"gtecsfield=Field"`
		GTEField      string `validate:"gtefield=Field"`
		GTField       string `validate:"gtfield=Field"`
		LTCSField     string `validate:"ltcsfield=Field"`
		LTECSField    string `validate:"ltecsfield=Field"`
		LTEField      string `validate:"ltefield=Field"`
		LTField       string `validate:"ltfield=Field"`
		NECSField     string `validate:"necsfield=Field"`
		NEField       string `validate:"nefield=Field"`
	}
	assert := assert.New(t)
	schema, err := utils.AssertParse(assert, Test{})
	assert.Nil(err)

	assert.JSONEq(`{
		"type": "object",
		"properties": {
			"eqcsfield": { "type": "string", "description": "Field Equals Another Field (relative)" },
			"eqfield": { "type": "string", "description": "Field Equals Another Field" },
			"fieldcontains": { "type": "string", "description": "Check the indicated characters are present in the Field" },
			"fieldexcludes": { "type": "string", "description": "Check the indicated characters are not present in the field" },
			"gtcsfield": { "type": "string", "description": "Field Greater Than Another Relative Field" },
			"gtecsfield": { "type": "string", "description": "Field Greater Than or Equal To Another Relative Field" },
			"gtefield": { "type": "string", "description": "Field Greater Than or Equal To Another Field" },
			"gtfield": { "type": "string", "description": "Field Greater Than Another Field" },
			"ltcsfield": { "type": "string", "description": "Less Than Another Relative Field" },
			"ltecsfield": { "type": "string", "description": "Less Than or Equal To Another Relative Field" },
			"ltefield": { "type": "string", "description": "Less Than or Equal To Another Field" },
			"ltfield": { "type": "string", "description": "Less Than Another Field" },
			"necsfield": { "type": "string", "description": "Field Does Not Equal Another Field (relative)" },
			"nefield": { "type": "string", "description": "Field Does Not Equal Another Field" }
		}`, schema)
}
