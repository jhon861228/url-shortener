package e2e

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"url-shortener/internal/shorturl/app"
	"url-shortener/internal/shorturl/core"
	"url-shortener/internal/shorturl/infrastructure"

	"github.com/golang/mock/gomock"

	"github.com/aws/aws-lambda-go/events"
)

func TestShortUrlE2E_Save(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Crear un mock de la interfaz ShortUrlRepository
	mockRepo := infrastructure.NewMockShortUrlRepository(ctrl)

	mockRepo.EXPECT().
		Save(gomock.Any()). // Esperamos cualquier objeto en esta llamada
		Return(nil)         // Y que devuelva nil (sin errores)

	shortURLService := app.NewShortUrlService(mockRepo)
	handler := app.NewShortUrlLambdaHandler(shortURLService)

	reqBody := core.Url{
		ID:          "123",
		UrlShorted:  "short.url/123",
		UrlOriginal: "http://example.com",
		CreatedAt:   "2024-01-01",
	}

	// Convertir el cuerpo de la solicitud a JSON
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

	handlerResp, _ := handler.HandleRequest(context.Background(), apiRequest)

	recorder.Result()

	if handlerResp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status %d but got %d", http.StatusCreated, handlerResp.StatusCode)
	}
}

func TestShortUrlE2E_Save_InvalidRequest(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Crear un mock de la interfaz ShortUrlRepository
	mockRepo := infrastructure.NewMockShortUrlRepository(ctrl)

	shortURLService := app.NewShortUrlService(mockRepo)
	handler := app.NewShortUrlLambdaHandler(shortURLService)

	// Convertir el cuerpo de la solicitud a JSON
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

	handlerResp, _ := handler.HandleRequest(context.Background(), apiRequest)

	recorder.Result()

	if handlerResp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status %d but got %d", http.StatusBadRequest, handlerResp.StatusCode)
	}
}

func TestShortUrlE2E_Save_MethodNotAllowed(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Crear un mock de la interfaz ShortUrlRepository
	mockRepo := infrastructure.NewMockShortUrlRepository(ctrl)

	shortURLService := app.NewShortUrlService(mockRepo)
	handler := app.NewShortUrlLambdaHandler(shortURLService)

	// Convertir el cuerpo de la solicitud a JSON
	reqBodyBytes, err := json.Marshal("{'field': 'error'}")
	if err != nil {
		t.Fatalf("Failed to marshal request body: %v", err)
	}

	req := httptest.NewRequest(http.MethodPut, "/short-url", bytes.NewBuffer(reqBodyBytes))
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	apiRequest := events.APIGatewayProxyRequest{
		HTTPMethod: http.MethodPut,
		Body:       string(reqBodyBytes),
	}

	handlerResp, _ := handler.HandleRequest(context.Background(), apiRequest)

	recorder.Result()

	if handlerResp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Expected status %d but got %d", http.StatusMethodNotAllowed, handlerResp.StatusCode)
	}
}

func TestShortUrlE2E_Save_Fail_Creation_Url(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Crear un mock de la interfaz ShortUrlRepository
	mockRepo := infrastructure.NewMockShortUrlRepository(ctrl)

	mockRepo.EXPECT().
		Save(gomock.Any()).                       // Esperamos cualquier objeto en esta llamada
		Return(fmt.Errorf("some error occurred")) // Y que devuelva nil (sin errores)

	shortURLService := app.NewShortUrlService(mockRepo)
	handler := app.NewShortUrlLambdaHandler(shortURLService)

	// Convertir el cuerpo de la solicitud a JSON

	reqBody := core.Url{
		ID:          "123",
		UrlShorted:  "short.url/123",
		UrlOriginal: "http://example.com",
		CreatedAt:   "2024-01-01",
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

	handlerResp, _ := handler.HandleRequest(context.Background(), apiRequest)

	recorder.Result()

	if handlerResp.StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected status %d but got %d", http.StatusInternalServerError, handlerResp.StatusCode)
	}
}

func TestShortUrlE2E_GetItem(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Crear un mock de la interfaz ShortUrlRepository
	mockRepo := infrastructure.NewMockShortUrlRepository(ctrl)

	mockUrl := &core.Url{
		ID:          "abc",
		UrlShorted:  "short.url/123",
		UrlOriginal: "http://example.com",
		CreatedAt:   "2024-01-01",
	}

	mockRepo.EXPECT().
		GetByID("abc").      // Esperamos cualquier objeto en esta llamada
		Return(mockUrl, nil) // Y que devuelva nil (sin errores)

	shortURLService := app.NewShortUrlService(mockRepo)
	handler := app.NewShortUrlLambdaHandler(shortURLService)

	req := httptest.NewRequest(http.MethodGet, "/short-url?id=abc", nil)
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	apiRequest := events.APIGatewayProxyRequest{
		HTTPMethod: http.MethodGet,
		QueryStringParameters: map[string]string{
			"id": "abc",
		},
	}

	handlerResp, _ := handler.HandleRequest(context.Background(), apiRequest)

	recorder.Result()

	if handlerResp.StatusCode != http.StatusOK {
		t.Errorf("Expected status %d but got %d", http.StatusCreated, handlerResp.StatusCode)
	}

	expectedBody, err := json.Marshal(mockUrl)
	if err != nil {
		t.Fatalf("Failed to marshal expected response body: %v", err)
	}

	if handlerResp.Body != string(expectedBody) { // Aquí necesitas comparar con la representación esperada del objeto mockUrl
		t.Errorf("Expected body '%s' but got '%s'", string(expectedBody), handlerResp.Body)
	}
}

func TestShortUrlE2E_GetItem_Fail(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Crear un mock de la interfaz ShortUrlRepository
	mockRepo := infrastructure.NewMockShortUrlRepository(ctrl)

	mockUrl := &core.Url{
		ID:          "abc",
		UrlShorted:  "short.url/123",
		UrlOriginal: "http://example.com",
		CreatedAt:   "2024-01-01",
	}

	mockRepo.EXPECT().
		GetByID("abc").                      // Esperamos cualquier objeto en esta llamada
		Return(mockUrl, fmt.Errorf("Error")) // Y que devuelva nil (sin errores)

	shortURLService := app.NewShortUrlService(mockRepo)
	handler := app.NewShortUrlLambdaHandler(shortURLService)

	req := httptest.NewRequest(http.MethodGet, "/short-url?id=abc", nil)
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	apiRequest := events.APIGatewayProxyRequest{
		HTTPMethod: http.MethodGet,
		QueryStringParameters: map[string]string{
			"id": "abc",
		},
	}

	handlerResp, _ := handler.HandleRequest(context.Background(), apiRequest)

	recorder.Result()

	if handlerResp.StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected status %d but got %d", http.StatusInternalServerError, handlerResp.StatusCode)
	}
}

func TestShortUrlE2E_GetItem_UrlNotFound(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Crear un mock de la interfaz ShortUrlRepository
	mockRepo := infrastructure.NewMockShortUrlRepository(ctrl)

	mockRepo.EXPECT().
		GetByID("abc").  // Esperamos cualquier objeto en esta llamada
		Return(nil, nil) // Y que devuelva nil (sin errores)

	shortURLService := app.NewShortUrlService(mockRepo)
	handler := app.NewShortUrlLambdaHandler(shortURLService)

	req := httptest.NewRequest(http.MethodGet, "/short-url?id=abc", nil)
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	apiRequest := events.APIGatewayProxyRequest{
		HTTPMethod: http.MethodGet,
		QueryStringParameters: map[string]string{
			"id": "abc",
		},
	}

	handlerResp, _ := handler.HandleRequest(context.Background(), apiRequest)

	recorder.Result()

	if handlerResp.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status %d but got %d", http.StatusNotFound, handlerResp.StatusCode)
	}
}
