package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type User struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

func handlerWrapper(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Declare schema
		return h(c)
	}
}

func GetUser(c echo.Context) error {
	var user User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}
	if err := c.Validate(&user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Validation failed", "details": err.Error()})
	}
	return c.JSON(http.StatusOK, user)
}

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/users", handlerWrapper(GetUser))

	e.Logger.Fatal(e.Start(":1323"))
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}