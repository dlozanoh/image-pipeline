package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	rekognition "github.com/aws/aws-sdk-go-v2/service/rekognition"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
)

var rekogClient *rekognition.Client

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	rekogClient = rekognition.NewFromConfig(cfg)
}

type DetectLabelsOutput struct {
	Labels []string `json:"labels"`
}

// Handler processes an S3 event and detects labels in the image
func Handler(ctx context.Context, s3Event events.S3Event) (DetectLabelsOutput, error) {
	if len(s3Event.Records) == 0 {
		return DetectLabelsOutput{}, fmt.Errorf("no records found in event")
	}

	record := s3Event.Records[0]
	bucket := record.S3.Bucket.Name
	key := record.S3.Object.Key

	key = strings.ReplaceAll(key, "+", " ")
	key = strings.ReplaceAll(key, "%20", " ")

	log.Printf("Detecting labels in image: s3://%s/%s", bucket, key)

	input := &rekognition.DetectLabelsInput{
		Image: &types.Image{
			S3Object: &types.S3Object{
				Bucket: &bucket,
				Name:   &key,
			},
		},
		MaxLabels:     aws.Int32(10),
		MinConfidence: aws.Float32(70.0),
	}

	result, err := rekogClient.DetectLabels(ctx, input)
	if err != nil {
		return DetectLabelsOutput{}, fmt.Errorf("failed to detect labels: %w", err)
	}

	var labels []string
	for _, label := range result.Labels {
		labels = append(labels, aws.ToString(label.Name))
	}

	log.Printf("Labels detected: %v", labels)

	return DetectLabelsOutput{
		Labels: labels,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
