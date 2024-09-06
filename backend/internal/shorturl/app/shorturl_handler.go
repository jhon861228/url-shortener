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

func (h *ShortUrlLambdaHandler) HandleGetRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return h.handleGetUrl(request)
}

func (h *ShortUrlLambdaHandler) HandlePostRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return h.handleCreateUrl(request)
}

func (h *ShortUrlLambdaHandler) handleCreateUrl(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var body core.UrlRequest

	json.Unmarshal([]byte(request.Body), &body)

	if err := body.ValidateUrlRequest(); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "Invalid request body",
		}, nil
	}

	url, err := h.ShortUrlService.CreateUrl(&body)
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Failed to create Url",
		}, nil
	}

	bodyResponse, _ := json.Marshal(core.UrlResponse{UrlShorted: url.UrlShorted})
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
		Body:       string(bodyResponse),
	}, nil
}

func (h *ShortUrlLambdaHandler) handleGetUrl(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id := request.PathParameters["urlId"]
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

	if _, err := json.Marshal(Url); err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Failed to marshal response",
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusFound,
		Body:       "",
		Headers:    map[string]string{"Location": Url.UrlOriginal},
	}, nil
}
