package main

import (
	shared "url-shortener/internal/shared/app"
	sharedConfig "url-shortener/internal/shared/config"
	shortUrl "url-shortener/internal/shorturl/app"
	"url-shortener/internal/shorturl/infrastructure"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {

	cfg := sharedConfig.LoadConfig()

	// Inyección de dependencias
	shortURLRepo := infrastructure.NewDynamoDBShortUrlRepository(cfg.DB, cfg.TableName)
	shortURLService := shortUrl.NewShortUrlService(shortURLRepo, cfg.UrlOriginalIndex, cfg.UrlOriginalIndexField)
	shortURLHandler := shortUrl.NewShortUrlLambdaHandler(shortURLService)

	// Routers
	router := shared.NewLambdaRouter()

	router.AddRouteWithMethod("/short-url", "POST", shortURLHandler.HandlePostRequest)

	router.AddRouteWithPathParam("/{urlId}", "GET", shortURLHandler.HandleGetRequest)

	lambda.Start(router.HandleRequest)
}
