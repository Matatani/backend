package uc_upload

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	tusd "github.com/tus/tusd/pkg/handler"
	"github.com/tus/tusd/pkg/s3store"
	"log"
	"net/http"
	"www.github.com/Maevlava/Matatani/backend/internal/config"
	"www.github.com/Maevlava/Matatani/backend/internal/middleware"
	"www.github.com/Maevlava/Matatani/backend/internal/predictor_service"
)

const BASE_PATH = "/upload/"

func TusHandler(cfg *config.APIConfig, predictorClient predictor_service.PredictorClient) (*tusd.Handler, error) {

	awsCfg := &aws.Config{
		Endpoint:         aws.String(cfg.S3Endpoint),
		Region:           aws.String(cfg.AWSRegion),
		S3ForcePathStyle: aws.Bool(true),
		Credentials:      credentials.NewStaticCredentials(cfg.AWSAccessKey, cfg.AWSSecretKey, ""),
		DisableSSL:       aws.Bool(true),
	}

	sess := session.Must(session.NewSession(awsCfg))
	svc := s3.New(sess)

	store := s3store.New(cfg.S3Bucket, svc)
	composer := tusd.NewStoreComposer()
	store.UseIn(composer)

	handler, err := tusd.NewHandler(tusd.Config{
		BasePath:              BASE_PATH,
		StoreComposer:         composer,
		NotifyCompleteUploads: true,
	})
	if err != nil {
		return nil, err
	}

	// listen for completed uploads
	go func() {
		for {
			event := <-handler.CompleteUploads
			handleUploadFinished(event, predictorClient, cfg)
		}
	}()

	return handler, nil
}

func handleUploadFinished(
	event tusd.HookEvent,
	predictorClient predictor_service.PredictorClient,
	cfg *config.APIConfig,
) error {

	s3Key, ok := event.Upload.Storage["Key"]
	if !ok {
		log.Println("S3 key not found in upload event")
		return nil
	}
	log.Printf("S3 key received: %s\n", s3Key)

	response, err := predictorClient.PredictImage(context.Background(), &predictor_service.PredictImageRequest{
		Bucket: cfg.S3Bucket,
		Key:    s3Key,
	})
	if err != nil {
		log.Printf("Error calling PredictImage: %v\n", err)
		return err
	}

	log.Printf("Prediction response: %v\n", response.ClassName)
	return nil
}
func RegisterUploadRoutes(router *http.ServeMux, tushandler *tusd.Handler) {

	// Middleware : CORS
	finalHandler := middleware.EnableCORS(tushandler)

	router.Handle(BASE_PATH, http.StripPrefix(BASE_PATH, finalHandler))
}
