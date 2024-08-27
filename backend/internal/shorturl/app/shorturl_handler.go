package app

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"url-shortener/internal/shorturl/core"

	"github.com/aws/aws-lambda-go/events"
)

type ShortUrlLambdaHandler struct {
	ShortUrlService core.ShortUrlService
}

func NewShortUrlLambdaHandler(service core.ShortUrlService) *ShortUrlLambdaHandler {
	return &ShortUrlLambdaHandler{ShortUrlService: service}
}

func (h *ShortUrlLambdaHandler) HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch request.HTTPMethod {
	case http.MethodPost:
		return h.handleCreateUrl(request)
	case http.MethodGet:
		return h.handleGetUrl(request)
	default:
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusMethodNotAllowed,
			Body:       "Method not allowed",
		}, nil
	}
}

func (h *ShortUrlLambdaHandler) handleCreateUrl(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var Url core.Url

	json.Unmarshal([]byte(request.Body), &Url)

	if err := Url.Validate(); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "Invalid request body",
		}, nil
	}

	if err := h.ShortUrlService.CreateUrl(&Url); err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Failed to create Url",
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
		Body:       "Url created successfully",
	}, nil
}

func (h *ShortUrlLambdaHandler) handleGetUrl(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id := request.QueryStringParameters["id"]
	Url, err := h.ShortUrlService.GetUrl(id)
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Failed to retrieve Url",
		}, nil
	}

	if Url == nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusNotFound,
			Body:       "Url not found",
		}, nil
	}

	body, err := json.Marshal(Url)
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Failed to marshal response",
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(body),
	}, nil
}
