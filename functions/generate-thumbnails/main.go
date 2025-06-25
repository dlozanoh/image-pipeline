package main

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"strings"

	_ "image/png"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/nfnt/resize"
)

var s3Client *s3.Client

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("Failed to load AWS config: %v", err)
	}
	s3Client = s3.NewFromConfig(cfg)
}

func Handler(ctx context.Context, s3Event events.S3Event) error {
	if len(s3Event.Records) == 0 {
		return fmt.Errorf("no S3 records in event")
	}

	record := s3Event.Records[0]
	bucket := record.S3.Bucket.Name
	key := record.S3.Object.Key
	key = strings.ReplaceAll(key, "+", " ")

	// Get image from S3
	getObj, err := s3Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return fmt.Errorf("failed to fetch object from S3: %w", err)
	}
	defer getObj.Body.Close()

	// Decode image
	img, _, err := image.Decode(getObj.Body)
	if err != nil {
		return fmt.Errorf("failed to decode image: %w", err)
	}

	// Resize image
	thumbnail := resize.Thumbnail(200, 200, img, resize.Lanczos3)

	// Encode thumbnail to buffer
	buf := new(bytes.Buffer)
	if err := jpeg.Encode(buf, thumbnail, nil); err != nil {
		return fmt.Errorf("failed to encode thumbnail: %w", err)
	}

	// Upload thumbnail to S3 (same bucket, different prefix)
	thumbKey := "thumbnails/" + key
	_, err = s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(thumbKey),
		Body:        bytes.NewReader(buf.Bytes()),
		ContentType: aws.String("image/jpeg"),
	})
	if err != nil {
		return fmt.Errorf("failed to upload thumbnail: %w", err)
	}

	log.Printf("Thumbnail created: s3://%s/%s", bucket, thumbKey)
	return nil
}

func main() {
	lambda.Start(Handler)
}
