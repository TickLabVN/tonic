package tonic

import (
	"fmt"
	"reflect"
	"regexp"
	"slices"
	"strings"
)

const (
	Get     = "GET"
	Post    = "POST"
	Put     = "PUT"
	Delete  = "DELETE"
	Patch   = "PATCH"
	Options = "OPTIONS"
	Head    = "HEAD"
)

const (
	pathParam  = "path"
	queryParam = "query"
)

type Route struct {
	Method  string
	Url     string
	Tags    []string
	Schema  *RouteSchema
	HandlerRegister func(path string)
}

type RouteSchema struct {
	Summary     string
	Description string
	Querystring any
	Params      any
	Body        any
	Response    map[int]any
}

var apiSpec map[string]any

func CreateRoutes(basePath string, routes []Route) {
	if !isInit {
		panic("Tonic must be initialized first!")
	}

	apiSpecPaths, _ := apiSpec["paths"].(map[string]any)

	for routeIdx := range routes {
		route := &routes[routeIdx]
		route.Url = normalizePath(route.Url)

		// Register the route handler
		route.HandlerRegister(route.Url)

		apiPath := toSwaggerAPIPath(fmt.Sprintf("%s%s", basePath, route.Url))

		pathSpec, apiPathExisted := apiSpecPaths[apiPath]
		if !apiPathExisted {
			pathSpec = make(map[string]any)
			apiSpecPaths[apiPath] = pathSpec
		}
		pathSpec.(map[string]any)[strings.ToLower(route.Method)] = buildHandlerSpec(route)
	}
}

func buildHandlerSpec(route *Route) map[string]any {
	handlerSpec := make(map[string]any)
	if route.Schema == nil {
		return handlerSpec
	}

	if route.Schema.Summary != "" {
		handlerSpec["summary"] = route.Schema.Summary
	}
	if route.Schema.Description != "" {
		handlerSpec["description"] = route.Schema.Description
	}
	if route.Tags != nil {
		handlerSpec["tags"] = route.Tags
	}

	var paramsSpec []map[string]any
	if route.Schema.Params != nil {
		paramsSpec = buildParamSpecs(pathParam, &route.Schema.Params)
	}
	if route.Schema.Querystring != nil {
		paramsSpec = append(paramsSpec, buildParamSpecs(queryParam, &route.Schema.Querystring)...)
	}
	if len(paramsSpec) > 0 {
		handlerSpec["parameters"] = paramsSpec
	}

	if route.Schema.Body != nil {
		handlerSpec["requestBody"] = map[string]any{
			"content": map[string]any{
				"application/json": map[string]any{
					"schema": ToSwaggerSchema(reflect.TypeOf(route.Schema.Body)),
				},
			},
			"required": true,
		}
	}

	if route.Schema.Response != nil {
		handlerSpec["responses"] = map[int]any{}
		for status, response := range route.Schema.Response {
			respType := reflect.TypeOf(response)
			handlerSpec["responses"].(map[int]any)[status] = map[string]any{
				"content": map[string]any{
					"application/json": map[string]any{
						"schema": ToSwaggerSchema(respType),
					},
				},
				"description": "Default description",
			}
		}
	}

	return handlerSpec
}

func buildParamSpecs(paramType string, params *any) []map[string]any {
	paramSpecs := []map[string]any{}
	t := reflect.TypeOf(*params)
	paramObjSchema := ToSwaggerSchema(t)

	for propName, propSchema := range paramObjSchema["properties"].(map[string]any) {
		paramSchema := map[string]any{
			"in":       paramType,
			"name":     propName,
			"required": slices.Contains(paramObjSchema["required"].([]string), propName),
			"schema":   propSchema,
		}
		paramSpecs = append(paramSpecs, paramSchema)
	}

	return paramSpecs
}

func toSwaggerAPIPath(path string) string {
	modifiedPath := regexp.MustCompile(`/:(\w+)`).ReplaceAllStringFunc(path, func(match string) string {
		return fmt.Sprintf("/{%s}", match[2:])
	})

	return modifiedPath
}

func normalizePath(path string) string {
	if path == "/" {
		return ""
	} else if len(path) > 0 && path[0] != '/' {
		return fmt.Sprintf("/%s", path)
	}
	return path
}
