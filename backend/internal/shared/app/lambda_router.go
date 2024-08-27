package app

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type LambdaRouter struct {
	routes map[string]func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
}

func NewLambdaRouter() *LambdaRouter {
	return &LambdaRouter{
		routes: make(map[string]func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)),
	}
}

func (r *LambdaRouter) AddRoute(path string, handler func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)) {
	r.routes[path] = handler
}

func (r *LambdaRouter) Handle(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	handler, ok := r.routes[request.Resource]
	if !ok {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusNotFound,
			Body:       "Resource not found",
		}, nil
	}

	return handler(ctx, request)
}
