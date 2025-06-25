package json_test

import (
	"fmt"
	"testing"

	json2 "encoding/json"

	"github.com/TickLabVN/tonic/core/docs"
	"github.com/TickLabVN/tonic/core/json"
)

// generic test suite
type testSuite struct {
	name  string
	input interface{}
}

func (ts *testSuite) run(t *testing.T) {
	t.Run(ts.name, func(t *testing.T) {
		t.Logf("Running test %s", ts.name)
		bytes, err := json.MarshalInline(ts.input)
		if err != nil {
			t.Error(err)
		}

		fmt.Println(string(bytes))
	})
}

// go test -v -run TestMalsalInlineWithOneFieldIsNill github.com/TickLabVN/tonic/core/json
func TestMalsalInlineWithOneFieldIsNill(t *testing.T) {
	suite := &testSuite{
		name: "TestMalsalInlineWithOneFieldIsNill",
		input: struct {
			Name string `json:"name"`
			Age  *int   `json:"age,omitempty"`
		}{
			Name: "John",
			Age:  nil,
		},
	}
	suite.run(t)
}

// go test -v -run TestMalsalInlineWithOneFieldIsNilAndNotHaveJsonTag github.com/TickLabVN/tonic/core/json
func TestMalsalInlineWithOneFieldIsNilAndNotHaveJsonTag(t *testing.T) {
	suite := &testSuite{
		name: "TestMalsalInlineWithOneFieldIsNilAndNotHaveJsonTag",
		input: struct {
			Name string
			Age  *int
		}{
			Name: "John",
			Age:  nil,
		},
	}
	suite.run(t)
}

func TestMarshalInlineInt32(t *testing.T) {
	suite := &testSuite{
		name:  "TestMarshalInlineInt32",
		input: 10,
	}

	suite.run(t)
}

func TestMashalInlineMapString(t *testing.T) {
	suite := &testSuite{
		name: "TestMashalInlineMapString",
		input: map[string]int{
			"name": 10,
			"age":  20,
		},
	}

	jsonData, _ := json2.Marshal(suite.input)
	fmt.Println(string(jsonData))

	suite.run(t)
}

func TestMarshalInlineStruct(t *testing.T) {
	suite := &testSuite{
		name: "TestMarshalInlineStruct",
		input: struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}{
			Name: "John",
			Age:  20,
		},
	}

	suite.run(t)
}

// go test -v -run TestMarshalMap github.com/TickLabVN/tonic/core/json
func TestMarshalMap(t *testing.T) {
	type Map map[string]struct {
		Name   string `json:"name"`
		School *struct {
			ClassName string `json:"class_name"`
			Number    int    `json:"number"`
		} `json:",inline,omitempty"`
	}

	suite := &testSuite{
		name: "TestMarshalMap",
		input: &Map{
			"12A1": {
				Name: "John",
				School: &struct {
					ClassName string `json:"class_name"`
					Number    int    `json:"number"`
				}{
					ClassName: "School 1",
					Number:    1,
				},
			},
		},
	}
	suite.run(t)
}

// go test -v -run TestMarshalSpec github.com/TickLabVN/tonic/core/json
func TestMarshalSpec(t *testing.T) {
	paths := docs.NewPaths()
	paths.AddPath("/categories", &docs.PathItem{
		Get: &docs.Operation{
			Description: "Returns all categories from the system that the user has access to",
			Parameters: []*docs.ParameterOrReference{
				{
					Reference: &docs.Reference{
						Description: "A parameter",
						Summary:     "A parameter",
						Ref:         "#/components/parameters/limitParam",
					},
				},
				{
					Parameter: &docs.Parameter{
						Name: "skip",
						In:   "query",
					},
				},
			},
		},
	})

	suite := &testSuite{
		name:  "TestMarshalSpec",
		input: paths,
	}

	suite.run(t)
}
