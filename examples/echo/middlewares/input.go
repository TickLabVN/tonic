package middlewares

import (
	"net/http"

	"echo_example/utils"

	"github.com/labstack/echo/v4"
)

func Bind[D any](next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var data D
		if err := c.Bind(&data); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
		}
		if err := c.Validate(&data); err != nil {
			return c.JSON(http.StatusBadRequest, utils.ValidateErrorMapping(err))
		}
		c.Set("data", data)
		return next(c)
	}
}
