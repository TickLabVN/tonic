package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	pkg "github.com/TickLabVN/tonic"
	"github.com/TickLabVN/tonic/schema"
	"github.com/flowchartsman/swaggerui"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	pkg.Init(
		pkg.WithInfo(&schema.Info{
			Title:       "Simple API",
			Version:     "0.0.3",
			Description: "This is a simple API for goem template",
		}),
		pkg.WithServers([]*schema.Server{
			{
				URL:         "http://localhost:{port}",
				Description: "Development server",
				Variables: map[string]schema.ServerVariableObject{
					"port": {
						Default:     "12345",
						Description: "Port number",
					},
				},
			},
			{
				URL:         "https://api.example.com",
				Description: "Production server",
			},
		}))

	// Add a path

	pkg.SetPath(map[string]*schema.Path{
		"/book": {
			Get: &schema.Operation{
				Summary: "Get a list of books",
				Tags:    []string{"book"},
			},
		},
	})

	s := pkg.GetSpec()
	fmt.Println("spec: ", pkg.GetSpec())

	specBytes, _ := json.Marshal(s)
	e.GET("/docs/*", echo.WrapHandler(http.StripPrefix("/docs", swaggerui.Handler(specBytes))))
	// curl http://localhost:12345/swagger/index.html
	e.Logger.Fatal(e.Start(":12345"))
}
