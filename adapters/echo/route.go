package echoAdapter

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/TickLabVN/tonic/core/docs"
	"github.com/TickLabVN/tonic/core/utils"
	"github.com/labstack/echo/v4"
)

type BindingOptions struct {
	Param  bool
	Query  bool
	Header bool
	Body   bool
}

func getParsingOptions(t reflect.Type) BindingOptions {
	opts := BindingOptions{}
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() == reflect.Struct {
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			opts.Param = opts.Param || field.Tag.Get("param") != ""
			opts.Query = opts.Query || field.Tag.Get("query") != ""
			opts.Header = opts.Header || field.Tag.Get("header") != ""
			opts.Body = opts.Body || field.Tag.Get("json") != ""
		}
	}
	return opts
}

func AddRoute[D any, R any](spec *docs.OpenApi, route *echo.Route, opts ...docs.OperationObject) {
	input, resp := reflect.TypeOf(new(D)), reflect.TypeOf(new(R))
	opId := fmt.Sprintf("%s_%s", route.Method, route.Name)
	opId = strings.ReplaceAll(opId, ".", "_")
	op := docs.OperationObject{
		OperationId: opId,
	}
	schemaBasePath := utils.GetSchemaPath(input)

	parsingOpts := getParsingOptions(input)
	if parsingOpts.Param {
		schema, err := spec.Components.AddSchema(input, "param", "validate")
		if err != nil {
			fmt.Printf("Error adding param schema: %v\n", err)
			return
		}
		op.AddParameter("path", schema, schemaBasePath+"_param")
	}
	if parsingOpts.Query {
		schema, err := spec.Components.AddSchema(input, "query", "validate")
		if err != nil {
			fmt.Printf("Error adding query schema: %v\n", err)
			return
		}
		op.AddParameter("query", schema, schemaBasePath+"_query")
	}
	if parsingOpts.Header {
		schema, err := spec.Components.AddSchema(input, "header", "validate")
		if err != nil {
			fmt.Printf("Error adding header schema: %v\n", err)
			return
		}
		op.AddParameter("header", schema, schemaBasePath+"_header")
	}
	if parsingOpts.Body {
		_, err := spec.Components.AddSchema(input, "json", "validate")
		if err != nil {
			fmt.Printf("Error adding body schema: %v\n", err)
			return
		}
		op.RequestBody = &docs.RequestBodyOrReference{
			RequestBodyObject: &docs.RequestBodyObject{
				Content: map[string]docs.MediaTypeObject{
					"application/json": {
						Schema: &docs.SchemaOrReference{
							ReferenceObject: &docs.ReferenceObject{
								Ref: schemaBasePath + "_json",
							},
						},
					},
				},
			},
		}
	}
	_, err := spec.Components.AddSchema(resp, "json", "validate")
	if err != nil {
		fmt.Printf("Error adding schema: %v\n", err)
		return
	}

	op = utils.MergeStructs(op, docs.OperationObject{
		Responses: map[string]docs.ResponseOrReference{
			"200": {
				ResponseObject: &docs.ResponseObject{
					Content: map[string]docs.MediaTypeObject{
						"application/json": {
							Schema: &docs.SchemaOrReference{
								ReferenceObject: &docs.ReferenceObject{
									Ref: utils.GetSchemaPath(resp) + "_json",
								},
							},
						},
					},
				},
			},
		},
	})
	op = utils.MergeStructs(append([]docs.OperationObject{op}, opts...)...)

	if spec.Paths == nil {
		spec.Paths = make(docs.Paths)
	}
	pathItem := docs.PathItemObject{}
	switch route.Method {
	case echo.GET:
		pathItem.Get = &op
	case echo.POST:
		pathItem.Post = &op
	case echo.PUT:
		pathItem.Put = &op
	case echo.PATCH:
		pathItem.Patch = &op
	case echo.DELETE:
		pathItem.Delete = &op
	case echo.OPTIONS:
		pathItem.Options = &op
	case echo.HEAD:
		pathItem.Head = &op
	default:
		fmt.Printf("Unsupported HTTP method: %s\n", route.Method)
	}

	routePath := strings.TrimSuffix(route.Path, "/")
	routePath = RE.ReplaceAllString(routePath, `{$1}`)
	spec.Paths.Update(routePath, pathItem)
}

var RE = regexp.MustCompile(`:([a-zA-Z0-9_]+)`)
