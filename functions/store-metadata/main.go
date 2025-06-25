package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

var (
	dynamoClient *dynamodb.Client
	tableName    = os.Getenv("TABLE_NAME")
)

type MetadataInput struct {
	ImageKey string   `json:"image_key"`
	Labels   []string `json:"labels"`
}

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	dynamoClient = dynamodb.NewFromConfig(cfg)
}

func Handler(ctx context.Context, input MetadataInput) (string, error) {
	if input.ImageKey == "" || len(input.Labels) == 0 {
		return "", fmt.Errorf("invalid input: image_key and labels are required")
	}

	labelsJson, err := json.Marshal(input.Labels)
	if err != nil {
		return "", fmt.Errorf("failed to encode labels: %w", err)
	}

	_, err = dynamoClient.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: map[string]types.AttributeValue{
			"image_key": &types.AttributeValueMemberS{Value: input.ImageKey},
			"labels":    &types.AttributeValueMemberS{Value: string(labelsJson)},
		},
	})
	if err != nil {
		return "", fmt.Errorf("failed to store metadata: %w", err)
	}

	log.Printf("Metadata stored for image: %s", input.ImageKey)
	return "Metadata stored successfully", nil
}

func main() {
	lambda.Start(Handler)
}
