package echoAdapter

import (
	"fmt"

	"github.com/TickLabVN/tonic/core"
	"github.com/TickLabVN/tonic/core/docs"
	"github.com/labstack/echo/v4"
)

func UIHandle(e *echo.Echo, spec *docs.OpenApi, path string) {
	swaggerPath := fmt.Sprintf("%s.json", path)
	e.GET(swaggerPath, echo.WrapHandler(core.JsonHttpHandler(spec)))
	e.GET(fmt.Sprintf("%s/*", path), echo.WrapHandler(core.SwaggerUIHandler(swaggerPath)))
}
