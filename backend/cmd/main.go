package main

import (
	"context"
	shared "url-shortener/internal/shared/app"
	shortUrl "url-shortener/internal/shorturl/app"
	"url-shortener/internal/shorturl/infrastructure"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func main() {

	cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
		o.Region = "us-east-1"
		return nil
	})
	if err != nil {
		panic(err)
	}
	db := dynamodb.NewFromConfig(cfg)

	// Inyección de dependencias
	shortURLRepo := infrastructure.NewDynamoDBShortUrlRepository(db, "shorturl-table")
	shortURLService := shortUrl.NewShortUrlService(shortURLRepo)
	shortURLHandler := shortUrl.NewShortUrlLambdaHandler(shortURLService)

	// Routers
	router := shared.NewLambdaRouter()

	router.AddRouteWithMethod("/short-url", "POST", shortURLHandler.HandlePostRequest)

	router.AddRouteWithPathParam("/{urlId}", "GET", shortURLHandler.HandleGetRequest)

	lambda.Start(router.HandleRequest)
}
