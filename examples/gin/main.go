package main

import (
	m "gin_example/middlewares"
	"net/http"

	ginAdapter "github.com/TickLabVN/tonic/adapters/gin"
	"github.com/TickLabVN/tonic/core/docs"
	"github.com/gin-gonic/gin"
)

type GetUserRequest struct {
	ID string `uri:"id" binding:"required"`
}

type User struct {
	ID    string `json:"id" binding:"required"`
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

func Bind[D any, R any](c *gin.Context) (D, error) {
	var req D
	if err := c.ShouldBind(&req); err != nil {
		return req, err
	}
	return req, nil
}

func GetUserById(c *gin.Context) {
	var req GetUserRequest
	if err := c.BindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := User{ID: req.ID, Name: "John Doe", Email: "john.doe@example.com"}
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.ID = "3" // Simulate ID generation
	c.JSON(http.StatusCreated, user)
}

func main() {
	// gin.SetMode(gin.ReleaseMode) // Set Gin to release mode for production
	g := gin.Default()
	openApi := &docs.OpenApi{
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
	ug := g.Group("/users")
	ginAdapter.AddRoute[GetUserRequest, User](openApi, ug, ginAdapter.Route{
		Method:   http.MethodGet,
		Path:     "/:id",
		Handlers: []gin.HandlerFunc{m.Bind[GetUserRequest], GetUserById},
	})
	ginAdapter.AddRoute[User, User](openApi, ug, ginAdapter.Route{
		Method:   http.MethodPost,
		Path:     "/",
		Handlers: []gin.HandlerFunc{m.Bind[User], CreateUser},
	})

	g.Run(":1234")
}
