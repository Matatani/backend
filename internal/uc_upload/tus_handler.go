package uc_upload

import (
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
)

const BASE_PATH = "/upload/"

func TusHandler(cfg *config.APIConfig) (*tusd.Handler, error) {

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
			handleUploadFinished(event)
		}
	}()

	return handler, nil
}

func handleUploadFinished(event tusd.HookEvent) {
	log.Printf("TUS: Upload finished!")
	log.Printf("TUS: File ID: %s", event.Upload.ID)
	log.Printf("TUS: File Size: %d", event.Upload.Size)
	log.Printf("TUS: File Path on Disk: %s", event.Upload.Storage["Path"])

	// TODO: Orchestration Logic
}
func RegisterUploadRoutes(router *http.ServeMux, tushandler *tusd.Handler) {

	// Middleware : CORS
	finalHandler := middleware.EnableCORS(tushandler)

	router.Handle(BASE_PATH, http.StripPrefix(BASE_PATH, finalHandler))
}
