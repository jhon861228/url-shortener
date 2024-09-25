package e2e

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"url-shortener/internal/shorturl/app"
	"url-shortener/internal/shorturl/core"
	"url-shortener/internal/shorturl/infrastructure"

	"github.com/golang/mock/gomock"

	"github.com/aws/aws-lambda-go/events"
)

func TestShortUrlE2E_Save(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := infrastructure.NewMockShortUrlRepository(ctrl)

	mockRepo.EXPECT().
		Save(gomock.Any()).
		Return(nil)

	mockRepo.EXPECT().
		GetByField(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(nil, nil)

	shortURLService := app.NewShortUrlService(mockRepo, "arg1", "arg2")
	handler := app.NewShortUrlLambdaHandler(shortURLService)

	reqBody := core.UrlRequest{
		UrlOriginal: "http://example.com",
	}

	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatalf("Failed to marshal request body: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/short-url", bytes.NewBuffer(reqBodyBytes))
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	apiRequest := events.APIGatewayProxyRequest{
		HTTPMethod: http.MethodPost,
		Body:       string(reqBodyBytes),
	}

	handlerResp, _ := handler.HandlePostRequest(context.Background(), apiRequest)

	recorder.Result()

	if handlerResp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status %d but got %d", http.StatusCreated, handlerResp.StatusCode)
	}
}

func TestShortUrlE2E_Save_UrlOriginalExists(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUrls := []*core.Url{
		{
			ID:          "abc",
			UrlShorted:  "short.url/123",
			UrlOriginal: "http://example.com",
			CreatedAt:   time.Now().String(),
		},
	}

	mockRepo := infrastructure.NewMockShortUrlRepository(ctrl)

	mockRepo.EXPECT().
		GetByField(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(mockUrls, nil)

	shortURLService := app.NewShortUrlService(mockRepo, "arg1", "arg2")
	handler := app.NewShortUrlLambdaHandler(shortURLService)

	reqBody := core.UrlRequest{
		UrlOriginal: "http://example.com",
	}

	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatalf("Failed to marshal request body: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/short-url", bytes.NewBuffer(reqBodyBytes))
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	apiRequest := events.APIGatewayProxyRequest{
		HTTPMethod: http.MethodPost,
		Body:       string(reqBodyBytes),
	}

	handlerResp, _ := handler.HandlePostRequest(context.Background(), apiRequest)

	recorder.Result()

	if handlerResp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status %d but got %d", http.StatusCreated, handlerResp.StatusCode)
	}
}

func TestShortUrlE2E_Save_InvalidRequest(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := infrastructure.NewMockShortUrlRepository(ctrl)

	shortURLService := app.NewShortUrlService(mockRepo, "arg1", "arg2")
	handler := app.NewShortUrlLambdaHandler(shortURLService)

	reqBodyBytes, err := json.Marshal("{'field': 'error'}")
	if err != nil {
		t.Fatalf("Failed to marshal request body: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/short-url", bytes.NewBuffer(reqBodyBytes))
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	apiRequest := events.APIGatewayProxyRequest{
		HTTPMethod: http.MethodPost,
		Body:       string(reqBodyBytes),
	}

	handlerResp, _ := handler.HandlePostRequest(context.Background(), apiRequest)

	recorder.Result()

	if handlerResp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status %d but got %d", http.StatusBadRequest, handlerResp.StatusCode)
	}
}

func TestShortUrlE2E_Save_Fail_Creation_Url(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := infrastructure.NewMockShortUrlRepository(ctrl)

	mockRepo.EXPECT().
		GetByField(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(nil, nil)

	mockRepo.EXPECT().
		Save(gomock.Any()).
		Return(fmt.Errorf("some error occurred"))

	shortURLService := app.NewShortUrlService(mockRepo, "arg1", "arg2")
	handler := app.NewShortUrlLambdaHandler(shortURLService)

	reqBody := core.UrlRequest{
		UrlOriginal: "http://example.com",
	}

	reqBodyBytes, err := json.Marshal(reqBody)

	if err != nil {
		t.Fatalf("Failed to marshal request body: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/short-url", bytes.NewBuffer(reqBodyBytes))
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	apiRequest := events.APIGatewayProxyRequest{
		HTTPMethod: http.MethodPost,
		Body:       string(reqBodyBytes),
	}

	handlerResp, _ := handler.HandlePostRequest(context.Background(), apiRequest)

	recorder.Result()

	if handlerResp.StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected status %d but got %d", http.StatusInternalServerError, handlerResp.StatusCode)
	}
}

func TestShortUrlE2E_GetItem(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := infrastructure.NewMockShortUrlRepository(ctrl)

	mockUrl := &core.Url{
		ID:          "abc",
		UrlShorted:  "short.url/123",
		UrlOriginal: "http://example.com",
		CreatedAt:   time.Now().String(),
	}

	mockRepo.EXPECT().
		GetByID("abc").
		Return(mockUrl, nil)

	shortURLService := app.NewShortUrlService(mockRepo, "arg1", "arg2")
	handler := app.NewShortUrlLambdaHandler(shortURLService)

	req := httptest.NewRequest(http.MethodGet, "/abc", nil)
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	apiRequest := events.APIGatewayProxyRequest{
		HTTPMethod: http.MethodGet,
		PathParameters: map[string]string{
			"urlId": "abc",
		},
	}

	handlerResp, _ := handler.HandleGetRequest(context.Background(), apiRequest)

	recorder.Result()

	if handlerResp.StatusCode != http.StatusFound {
		t.Errorf("Expected status %d but got %d", http.StatusCreated, handlerResp.StatusCode)
	}

	_, err := json.Marshal(mockUrl)
	if err != nil {
		t.Fatalf("Failed to marshal expected response body: %v", err)
	}
}

func TestShortUrlE2E_GetItem_Fail(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := infrastructure.NewMockShortUrlRepository(ctrl)

	mockUrl := &core.Url{
		UrlShorted: "http://example.com",
	}

	mockRepo.EXPECT().
		GetByID("abc").
		Return(mockUrl, fmt.Errorf("Error"))

	shortURLService := app.NewShortUrlService(mockRepo, "arg1", "arg2")
	handler := app.NewShortUrlLambdaHandler(shortURLService)

	req := httptest.NewRequest(http.MethodGet, "/abc", nil)
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	apiRequest := events.APIGatewayProxyRequest{
		HTTPMethod: http.MethodGet,
		PathParameters: map[string]string{
			"urlId": "abc",
		},
	}

	handlerResp, _ := handler.HandleGetRequest(context.Background(), apiRequest)

	recorder.Result()

	if handlerResp.StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected status %d but got %d", http.StatusInternalServerError, handlerResp.StatusCode)
	}
}

func TestShortUrlE2E_GetItem_UrlNotFound(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := infrastructure.NewMockShortUrlRepository(ctrl)

	mockRepo.EXPECT().
		GetByID("abc").
		Return(nil, nil)

	shortURLService := app.NewShortUrlService(mockRepo, "arg1", "arg2")
	handler := app.NewShortUrlLambdaHandler(shortURLService)

	req := httptest.NewRequest(http.MethodGet, "/abc", nil)
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	apiRequest := events.APIGatewayProxyRequest{
		HTTPMethod: http.MethodGet,
		PathParameters: map[string]string{
			"urlId": "abc",
		},
	}

	handlerResp, _ := handler.HandleGetRequest(context.Background(), apiRequest)

	recorder.Result()

	if handlerResp.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status %d but got %d", http.StatusNotFound, handlerResp.StatusCode)
	}
}
