package source

import (
	"bytes"
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

type S3Source interface {
	UploadPhoto(photoData []byte, photoName string) (string, error)
	GetPresignedURL(photoName string) (string, error)
}

type s3Source struct {
	s3Client *s3.Client
}

func NewS3Source(regionStr string) S3Source {
	config, error := config.LoadDefaultConfig(context.TODO(), config.WithRegion(regionStr))
	if error != nil {
		log.Fatalf("NewS3Source - Failed to load AWS config: %v", error)
	}

	s3Client := s3.NewFromConfig(
		config,
	)

	return &s3Source{
		s3Client: s3Client,
	}
}

func (s *s3Source) UploadPhoto(photoData []byte, photoName string) (string, error) {
	output, error := s.s3Client.PutObject(
		context.TODO(),
		&s3.PutObjectInput{
			Bucket: aws.String("amzn-s3-unplash"),
			Key:    aws.String(photoName),
			Body:   bytes.NewReader(photoData),
		},
	)
	if error != nil {
		return "", error
	}
	return *output.ETag, nil
}

func (s *s3Source) GetPresignedURL(photoName string) (string, error) {
	return "", nil
}
