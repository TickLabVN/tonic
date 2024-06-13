package main

import (
	"net/http"

	pkg "github.com/TickLabVN/tonic"
	"github.com/TickLabVN/tonic/docs"
	"github.com/TickLabVN/tonic/json"
	"github.com/flowchartsman/swaggerui"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	pkg.Init(
		pkg.WithInfo(docs.Info{
			Title:       "Simple API",
			Version:     "0.0.3",
			Description: "This is a simple API for goem template",
		}),
		pkg.WithServers([]*docs.Server{
			{
				URL:         "http://localhost:{port}",
				Description: "Development server",
				Variables: map[string]docs.ServerVariable{
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
		}),
	)

<<<<<<< HEAD
=======
	// TODO: remove these lines, just testing
>>>>>>> 4245b3d (feat: add ref for objectts)
	s := pkg.GetSpec()
	s.Components = &docs.Components{
		Schemas: map[string]*docs.Schema{
			"GeneralError": {
				Type: "object",
				Properties: map[string]docs.Schema{
					"code": {
						Type:   "integer",
						Format: "int32",
					},
					"message": {
						Type: "string",
					},
				},
			},
			"Category": {
				Type: "object",
				Properties: map[string]docs.Schema{
					"id": {
						Type:   "integer",
						Format: "int64",
					},
					"name": {
						Type: "string",
					},
				},
			},
			"Tag": {
				Type: "object",
				Properties: map[string]docs.Schema{
					"id": {
						Type:   "integer",
						Format: "int64",
					},
					"name": {
						Type: "string",
					},
				},
			},
		},
		Parameters: map[string]*docs.Parameter{
			"skipParam": {
				Name:        "skip",
				In:          "query",
				Description: "number of items to skip",
				Required:    true,
				Schema: &docs.Schema{
					Type:    "integer",
					Integer: &docs.Integer{Format: "int32"},
				},
			},
			"limitParam": {
				Name:        "limit",
				In:          "query",
				Description: "max records to return",
				Required:    true,
				Schema: &docs.Schema{
					Type:    "integer",
					Integer: &docs.Integer{Format: "int32"},
				},
			},
		},
		Responses: map[string]*docs.Response{
			"NotFound": {
				Description: "Entity not found.",
			},
			"IllegalInput": {
				Description: "Illegal input for operation.",
			},
			"GeneralError": {
				Description: "General Error",
				Content: map[string]*docs.MediaType{
					"application/json": {
						Schema: &docs.Schema{
							Ref: &docs.Ref{Ref: "#/components/schemas/GeneralError"},
						},
					},
				},
			},
		},
		SecuritySchemes: map[string]*docs.SecurityScheme{
			"api_key": {
				Type: "apiKey",
				Name: "api_key",
				In:   "header",
			},
			"petstore_auth": {
				Type: "oauth2",
				Flows: &docs.OAuthFlows{
					Implicit: &docs.OAuthFlow{
						AuthorizationURL: "https://example.org/api/oauth/dialog",
						Scopes: map[string]string{
							"write:pets": "modify pets in your account",
							"read:pets":  "read your pets",
						},
					},
				},
			},
		},
		Examples: map[string]*docs.Example{
			"cat": {
				Summary: "An example of a cat",
				Value: map[string]interface{}{
					"name":    "Fluffy",
					"petType": "Cat",
					"color":   "White",
				},
			},
		},
	}
	s.Paths = &docs.Paths{
		"/categories": {
			Get: &docs.Operation{
				Description: "Returns all categories from the system that the user has access to",
				Parameters: []*docs.ParameterOrReference{
					{
						Ref: "#/components/parameters/limitParam",
					},
				},
			},
		},
		"/pets": {
			Get: &docs.Operation{
				Description: "Returns all pets from the system that the user has access to",
				Responses: &docs.Responses{
					"200": {
						Description: "A list of pets.",
						Content: map[string]*docs.MediaType{
							"application/json": {
								Schema: &docs.Schema{
									Type: "array",
									Array: &docs.Array{
										Items: &docs.Schema{
											Ref: &docs.Ref{Ref: "#/components/schemas/Category"},
										},
									},
								},
							},
						},
					},
					"default": {
						Description: "unexpected error",
						Content: map[string]*docs.MediaType{
							"application/json": {
								Schema: &docs.Schema{
									Ref: &docs.Ref{Ref: "#/components/schemas/GeneralError"},
								},
							},
						},
					},
				},
			},
		},
	}

	specBytes, _ := json.MalshalInline(s)
	e.GET("/docs/*", echo.WrapHandler(http.StripPrefix("/docs", swaggerui.Handler(specBytes))))
	// curl http://localhost:12345/docs/
	e.Logger.Fatal(e.Start(":12345"))
}
