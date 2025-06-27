package main

import (
	"net/http"

	m "echo_example/middlewares"
	"echo_example/utils"

	echoAdapter "github.com/TickLabVN/tonic/adapters/echo"
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

	openApiSpec := &docs.OpenApi{
		OpenAPI: "3.0.1",
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

	echoAdapter.AddRoute[GetUserRequest, User](
		openApiSpec,
		e.GET("/users/:id", GetUser, m.Bind[GetUserRequest]),
	)
	echoAdapter.AddRoute[User, User](
		openApiSpec,
		e.POST("/users", GetUser, m.Bind[User]),
	)

	echoAdapter.UIHandle(e, openApiSpec, "/docs")
	e.Logger.Fatal(e.Start(":1323"))
}
