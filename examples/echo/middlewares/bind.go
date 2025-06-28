package middlewares

import (
	"echo_example/utils"

	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

var DEFAULT_BINDER = &echo.DefaultBinder{}

func Bind[D any](next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var data D
		fmt.Printf("Binding data for type: %T\n", data)
		if err := c.Bind(&data); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
		}
		err := DEFAULT_BINDER.BindHeaders(c, &data)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid headers"})
		}
		if err := c.Validate(&data); err != nil {
			return c.JSON(http.StatusBadRequest, utils.ValidateErrorMapping(err))
		}
		c.Set("data", data)
		return next(c)
	}
}
