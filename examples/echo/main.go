package main

import (
	"net/http"

	"echo_example/middlewares"
	"echo_example/utils"

	echoAdapter "github.com/TickLabVN/tonic/adapters/echo"
	"github.com/TickLabVN/tonic/core/docs"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type GetUserRequest struct {
	ID     string `param:"id" validate:"required"`
	Name   string `query:"name"`
	ApiKey string `header:"x-api-key" validate:"required"`

	// Should be ignored in the schema if request is not a POST
	SampleJsonField string `json:"sampleJsonField" validate:"required"`
}

type User struct {
	ID    string `json:"id" validate:"required"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

func GetUser(c echo.Context) error {
	data := c.Get("data").(GetUserRequest)
	return c.JSON(http.StatusOK, User{
		ID:    data.ID,
		Name:  "John Doe",
		Email: "john.doe@example.com",
	})
}

func main() {
	e := echo.New()
	e.Validator = &utils.CustomValidator{Validator: validator.New()}

	openapi := &docs.OpenApi{
		OpenAPI: "3.0.1",
		Info: docs.InfoObject{
			Version: "1.0.0",
			Title:   "Echo Example API",
		},
	}

	echoAdapter.AddRoute[GetUserRequest, User](
		openapi,
		e.GET("/users/:id", GetUser, middlewares.Bind[GetUserRequest]),
	)
	echoAdapter.AddRoute[User, User](
		openapi,
		e.POST("/users", GetUser, middlewares.Bind[User]),
	)
	echoAdapter.UIHandle(e, openapi, "/docs")

	e.Logger.Fatal(e.Start(":1323"))
}
