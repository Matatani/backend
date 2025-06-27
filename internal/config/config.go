package config

import (
	"os"
)

type APIConfig struct {
	Host         string
	Port         string
	MLHost       string
	S3Bucket     string
	S3Endpoint   string
	AWSAccessKey string
	AWSSecretKey string
	AWSRegion    string
}

func Load() (*APIConfig, error) {
	host := os.Getenv("BE_HOST")
	port := os.Getenv("BE_PORT")
	mlHost := os.Getenv("ML_HOST")
	s3Bucket := os.Getenv("S3_BUCKET")
	s3Endpoint := os.Getenv("S3_ENDPOINT")
	awsAccessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	awsRegion := os.Getenv("AWS_REGION")
	return &APIConfig{
		Host:         host,
		Port:         port,
		MLHost:       mlHost,
		S3Bucket:     s3Bucket,
		S3Endpoint:   s3Endpoint,
		AWSAccessKey: awsAccessKey,
		AWSSecretKey: awsSecretKey,
		AWSRegion:    awsRegion,
	}, nil

}
