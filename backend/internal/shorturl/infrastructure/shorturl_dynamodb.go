package infrastructure

import (
	"context"
	"fmt"
	"url-shortener/internal/shorturl/core"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type DynamoDBShortUrlRepository struct {
	db    *dynamodb.Client
	table string
}

func NewDynamoDBShortUrlRepository(db *dynamodb.Client, table string) *DynamoDBShortUrlRepository {
	return &DynamoDBShortUrlRepository{
		db:    db,
		table: table,
	}
}

func (r *DynamoDBShortUrlRepository) Save(Url *core.Url) error {

	out, err := r.db.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(r.table),
		Item: map[string]types.AttributeValue{
			"id":          &types.AttributeValueMemberS{Value: Url.ID},
			"urlShorted":  &types.AttributeValueMemberS{Value: Url.UrlShorted},
			"urlOriginal": &types.AttributeValueMemberS{Value: Url.UrlOriginal},
			"createdAt":   &types.AttributeValueMemberS{Value: Url.CreatedAt},
		},
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(out.Attributes)

	return err
}

func (r *DynamoDBShortUrlRepository) GetByID(id string) (*core.Url, error) {
	result, err := r.db.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(r.table),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
	})

	if err != nil {
		panic(err)
	}

	if result.Item == nil {
		return nil, nil
	}

	Url := new(core.Url)

	err = attributevalue.UnmarshalMap(result.Item, Url)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal item: %v", err)
	}

	return Url, err
}

func (r *DynamoDBShortUrlRepository) GetByField(indexName string, fieldName string, fieldValue string) ([]*core.Url, error) {
	fmt.Println(r.table, indexName, fieldName, fieldValue)
	result, err := r.db.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:              aws.String(r.table),
		IndexName:              aws.String(indexName),
		KeyConditionExpression: aws.String("#field = :value"),
		ExpressionAttributeNames: map[string]string{
			"#field": fieldName,
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":value": &types.AttributeValueMemberS{Value: fieldValue},
		},
	})
	fmt.Println(result)
	fmt.Println(err)
	if err != nil {
		return nil, fmt.Errorf("failed to query items: %w", err)
	}
	var urls []*core.Url
	err = attributevalue.UnmarshalListOfMaps(result.Items, &urls)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal items: %v", err)
	}

	if len(urls) == 0 {
		return nil, nil
	}

	return urls, nil
}
