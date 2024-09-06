package app

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

type LambdaRoute struct {
	Method  string
	Handler func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
}

type LambdaRouter struct {
	routes map[string]LambdaRoute
}

func NewLambdaRouter() *LambdaRouter {
	return &LambdaRouter{
		routes: make(map[string]LambdaRoute),
	}
}

func (r *LambdaRouter) AddRoute(path string, handler func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)) {
	r.routes[path] = LambdaRoute{
		Method:  "",
		Handler: handler,
	}
}

func (r *LambdaRouter) AddRouteWithMethod(path string, method string, handler func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)) {
	r.routes[path] = LambdaRoute{
		Method:  method,
		Handler: handler,
	}
}

func (r *LambdaRouter) AddRouteWithPathParam(path string, method string, handler func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)) {
	r.routes[path] = LambdaRoute{
		Method:  method,
		Handler: handler,
	}
}

func (r *LambdaRouter) HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println(request.Path, request.HTTPMethod, request.Body, request.PathParameters)
	route, exists := r.routes[request.Path]

	if exists && request.HTTPMethod == route.Method {
		return route.Handler(ctx, request)
	}

	for routePattern, route := range r.routes {
		if match, params := matchPath(routePattern, request.Path); match && request.HTTPMethod == route.Method {
			request.PathParameters = params
			return route.Handler(ctx, request)
		}
	}

	if !exists {
		return events.APIGatewayProxyResponse{
			StatusCode: 404,
			Body:       "Not Found",
		}, nil
	}

	if route.Method != "" && route.Method != request.HTTPMethod {
		return events.APIGatewayProxyResponse{
			StatusCode: 405,
			Body:       "Method Not Allowed",
		}, nil
	}

	return route.Handler(ctx, request)
}

func matchPath(routePattern, requestPath string) (bool, map[string]string) {
	routeParts := strings.Split(routePattern, "/")
	requestParts := strings.Split(requestPath, "/")

	if len(routeParts) != len(requestParts) {
		return false, nil
	}

	params := make(map[string]string)
	for i := range routeParts {
		if strings.HasPrefix(routeParts[i], "{") && strings.HasSuffix(routeParts[i], "}") {
			paramName := routeParts[i][1 : len(routeParts[i])-1]
			params[paramName] = requestParts[i]
		} else if routeParts[i] != requestParts[i] {
			return false, nil
		}
	}
	return true, params
}
