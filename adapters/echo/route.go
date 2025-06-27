package echoAdapter

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/TickLabVN/tonic/core/docs"
	"github.com/TickLabVN/tonic/core/utils"
	"github.com/labstack/echo/v4"
)

func AddRoute[D any, R any](spec *docs.OpenApi, route *echo.Route, opts ...docs.OperationObject) {
	_, resp := reflect.TypeOf(new(D)), reflect.TypeOf(new(R))
	err := spec.Components.AddSchema(resp, "validate")
	if err != nil {
		fmt.Printf("Error adding schema: %v\n", err)
		return
	}

	op := utils.MergeStructs(opts...)
	opId := fmt.Sprintf("%s_%s", route.Method, route.Name)
	opId = strings.ReplaceAll(opId, ".", "_")
	op = utils.MergeStructs(op, docs.OperationObject{
		OperationId: opId,
		// Parameters:  docs.GetParametersFromType(input),
		Responses: map[string]docs.ResponseOrReference{
			"200":  {
				ResponseObject: &docs.ResponseObject{
					Content: map[string]docs.MediaTypeObject{
						"application/json": {
							Schema: &docs.SchemaOrReference{
								ReferenceObject: &docs.ReferenceObject{
									Ref: utils.GetSchemaPath(resp),
								},
							},
						},
					},
				},
			},
		},
	})

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
	fmt.Printf("Adding route: %s %s, %+v\n", route.Method, route.Path, op)
	spec.Paths.Update(route.Path, pathItem)
}
