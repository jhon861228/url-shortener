package config

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Config struct {
	Domain                string
	UrlOriginalIndex      string
	UrlOriginalIndexField string
	DB                    *dynamodb.Client
	TableName             string
}

func LoadConfig() *Config {
	domain := os.Getenv("DOMAIN")
	if domain == "" {
		log.Fatal("DOMAIN environment variable is required")
	}

	UrlOriginalIndex := os.Getenv("URLORIGINAL_INDEX")
	if UrlOriginalIndex == "" {
		log.Fatal("URLORIGINAL_INDEX environment variable is required")
	}

	UrlOriginalIndexField := os.Getenv("URLORIGINAL_INDEX_FIELD")
	if UrlOriginalIndexField == "" {
		log.Fatal("URLORIGINAL_INDEX_FIELD environment variable is required")
	}

	TableName := os.Getenv("TABLE_NAME")
	if TableName == "" {
		log.Fatal("TableName environment variable is required")
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
		o.Region = os.Getenv("APP_REGION")
		return nil
	})
	if err != nil {
		panic(err)
	}
	db := dynamodb.NewFromConfig(cfg)

	return &Config{
		Domain:                domain,
		UrlOriginalIndex:      UrlOriginalIndex,
		UrlOriginalIndexField: UrlOriginalIndexField,
		DB:                    db,
		TableName:             TableName,
	}
}
