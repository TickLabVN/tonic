package echoadapter

import (
	"reflect"

	"github.com/TickLabVN/tonic/core/docs"
	"github.com/TickLabVN/tonic/core/utils"
	"github.com/labstack/echo/v4"
)

func WrapGET[D any, R any](spec *docs.OpenApi, route *echo.Route, opts ...docs.OperationObject) {
	input, resp := reflect.TypeOf(new(D)), reflect.TypeOf(new(R))
	spec.Components.AddSchema(input)
	spec.Components.AddSchema(resp)

	getOperation := utils.MergeStructs(opts...)
	getOperation = utils.MergeStructs(getOperation, docs.OperationObject{
		OperationId: route.Name,
		// Parameters:  docs.GetParametersFromType(input),
		// Responses:   docs.GetResponsesFromType(resp),
	})
	spec.Paths.AddPath(route.Path, docs.PathItemObject{
		Get: &getOperation,
	})
}
