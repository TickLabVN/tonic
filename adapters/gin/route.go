package ginAdapter

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/TickLabVN/tonic/core/docs"
	"github.com/TickLabVN/tonic/core/utils"
	"github.com/gin-gonic/gin"
)

type Route struct {
	Method   string
	Path     string
	Handlers []gin.HandlerFunc
	opts     []docs.OperationObject
}

func AddRoute[D any, R any](spec *docs.OpenApi, g gin.IRoutes, route Route) {
	_, resp := reflect.TypeOf(new(D)), reflect.TypeOf(new(R))
	spec.Components.AddSchema(resp)

	var basePath string
	group, ok := g.(*gin.RouterGroup)
	if ok {
		basePath = group.BasePath()
	} else {
		engine, ok := g.(*gin.Engine)
		if ok {
			basePath = engine.BasePath()
		} else {
			panic("Invalid gin.IRoutes type, expected *gin.RouterGroup or *gin.Engine")
		}
	}

	baseOp := utils.MergeStructs(route.opts...)
	path := fmt.Sprintf("%s%s", basePath, route.Path)

	op := utils.MergeStructs(baseOp, docs.OperationObject{
		OperationId: fmt.Sprintf("%s_%s", route.Method, path),
		// Parameters:  docs.GetParametersFromType(input),
		Responses: map[string]docs.ResponseOrReference{
			"200": {
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
	case http.MethodGet:
		g.GET(path, route.Handlers...)
		pathItem.Get = &op
	case http.MethodPost:
		g.POST(path, route.Handlers...)
		pathItem.Post = &op
	case http.MethodPut:
		g.PUT(path, route.Handlers...)
		pathItem.Put = &op
	case http.MethodPatch:
		g.PATCH(path, route.Handlers...)
		pathItem.Patch = &op
	case http.MethodDelete:
		g.DELETE(path, route.Handlers...)
		pathItem.Delete = &op
	case http.MethodOptions:
		g.OPTIONS(path, route.Handlers...)
		pathItem.Options = &op
	case http.MethodHead:
		g.HEAD(path, route.Handlers...)
		pathItem.Head = &op
	default:
		fmt.Printf("Unsupported HTTP method: %s\n", route.Method)
	}
	spec.Paths.Update(path, pathItem)
}
