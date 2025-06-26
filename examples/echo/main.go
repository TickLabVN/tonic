package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	echoadapter "github.com/TickLabVN/tonic/adapters/echo"
	"github.com/TickLabVN/tonic/core/docs"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type GetUserRequest struct {
	ID string `param:"id" validate:"required"`
}

type User struct {
	ID    string `json:"id" validate:"required"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

func validateErrorMapping(err error) map[string]string {
	if err == nil {
		return nil
	}

	validationErrors := make(map[string]string)
	if _, ok := err.(*validator.InvalidValidationError); ok {
		validationErrors["error"] = "Invalid validation error"
		return validationErrors
	}

	for _, fieldErr := range err.(validator.ValidationErrors) {
		validationErrors[fieldErr.Field()] = fieldErr.Tag()
	}
	return validationErrors
}

func GetUser(c echo.Context) error {
	data := c.Get("data").(GetUserRequest)
	return c.JSON(http.StatusOK, User{
		ID:    data.ID,
		Name:  "John Doe",
		Email: "john.doe@example.com",
	})
}

func bindingMiddleware[D any](next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var data D
		if err := c.Bind(&data); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
		}
		if err := c.Validate(&data); err != nil {
			return c.JSON(http.StatusBadRequest, validateErrorMapping(err))
		}
		c.Set("data", data)
		return next(c)
	}
}

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	openApiSpec := &docs.OpenApi{
		OpenAPI: "3.1.1",
		Info: docs.InfoObject{
			Version: "1.0.0",
			Title:   "Echo Example API",
			Contact: &docs.ContactObject{
				Name:  "Author",
				URL:   "https://github.com/phucvinh57",
				Email: "npvinh0507@gmail.com",
			},
		},
	}
	echoadapter.[GetUserRequest, User](e.GET("/users/:id", GetUser, echoadapter.WithName("GetUser")), openApiSpec)

	specJson, _ := json.MarshalIndent(openApiSpec, "", "  ")
	fmt.Println("OpenAPI Specification:", string(specJson))
	e.Logger.Fatal(e.Start(":1323"))
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
