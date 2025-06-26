package schema_test

import (
	"testing"

	"github.com/TickLabVN/tonic/core"
	"github.com/stretchr/testify/assert"
)

func TestValidate_Strings(t *testing.T) {
	type Test struct {
		Alpha           string `json:"alpha" validate:"alpha"`
		Alphanum        string `json:"alphanum" validate:"alphanum"`
		AlphanumUnicode string `json:"alphanum_unicode" validate:"alphanum_unicode"`
		AlphaUnicode    string `json:"alpha_unicode" validate:"alpha_unicode"`
		ASCII           string `json:"ascii" validate:"ascii"`
		Boolean         string `json:"boolean" validate:"boolean"`
		Contains        string `json:"contains" validate:"contains=hello world"`
		ContainsAny     string `json:"contains_any" validate:"contains_any=hello world"`
		ContainsRune    string `json:"contains_rune" validate:"contains_rune=⌘"`
		EndsNotWith     string `json:"ends_not_with" validate:"ends_not_with=bye"`
		EndsWith        string `json:"ends_with" validate:"ends_with=bye"`
		Excludes        string `json:"excludes" validate:"excludes=foo"`
		ExcludesAll     string `json:"excludes_all" validate:"excludes_all=foo"`
		ExcludesRune    string `json:"excludes_rune" validate:"excludes_rune=⌘"`
		Number          string `json:"number" validate:"number"`
		Numeric         string `json:"numeric" validate:"numeric"`
		PrintASCII      string `json:"print_ascii" validate:"print_ascii"`
		StartsNotWith   string `json:"starts_not_with" validate:"starts_not_with=hello"`
		StartsWith      string `json:"starts_with" validate:"starts_with=bye"`
		Uppercase       string `json:"uppercase" validate:"uppercase"`
		Lowercase       string `json:"lowercase" validate:"lowercase"`
		Multibyte       string `json:"multibyte" validate:"multibyte"`
	}

	spec := core.Init()
	assert := assert.New(t)
	result, err := AssertParse(assert, spec, Test{})
	assert.Nil(err)

	assert.JSONEq(`{
		"type": "object",
		"properties": {
			"alpha": {"type": "string", "format": "alpha", "pattern": "^[a-zA-Z]*$"},
			"alphanum": {"type": "string", "format": "alphanum", "pattern": "^[a-zA-Z0-9]*$"},
			"alphanum_unicode": {"type": "string", "format": "alphanum_unicode", "pattern": "^[\\p{L}0-9]*$"},
			"alpha_unicode": {"type": "string", "format": "alpha_unicode", "pattern": "^[\\p{L}]*$"},
			"ascii": {"type": "string", "format": "ascii", "pattern": "^[\\x00-\\x7F]*$"},
			"boolean": {"type": "string", "format": "boolean", "enum": ["true", "false"]},
			"contains": {"type": "string", "description": "Contain 'hello world'", "pattern": ".*hello world.*"},
			"contains_any": {"type": "string", "description": "Contain any chars in 'hello world'", "pattern": ".*[hello world].*"},
			"contains_rune": {"type": "string", "description": "Contain '⌘'", "pattern": ".*⌘.*"},
			"ends_not_with": {"type": "string", "description": "Not end with 'bye'", "not": {"pattern": ".*bye$"}},
			"ends_with": {"type": "string", "description": "End with 'bye'", "pattern": ".*bye$"},
			"excludes": {"type": "string", "description": "Not contain 'foo'", "not": {"pattern": ".*foo.*"}},
			"excludes_all": {"type": "string", "description": "Not contain any chars in 'foo'", "not": {"pattern": ".*[foo].*"}},
			"excludes_rune": {"type": "string", "description": "Not contain '⌘'", "not": {"pattern": ".*⌘.*"}},
			"number": {"type": "string", "format": "number", "pattern": "^[0-9]*$"},
			"numeric": {"type": "string", "format": "numeric", "pattern": "^[0-9]*$"},
			"print_ascii": {"type": "string", "format": "print_ascii", "pattern": "^[\\x20-\\x7E]*$"},	
			"starts_not_with": {"type": "string", "description": "Not start with 'hello'", "not": {"pattern": "^hello.*"}},
			"starts_with": {"type": "string", "description": "Start with 'bye'", "pattern": "^bye.*"},
			"uppercase": {"type": "string", "format": "uppercase", "pattern": "^[A-Z]*$"},
			"lowercase": {"type": "string", "format": "lowercase", "pattern": "^[a-z]*$"},
			"multibyte": {"type": "string", "format": "multibyte", "pattern": "^[\\p{L}0-9]*$"}
		}
	}`, result)
}
