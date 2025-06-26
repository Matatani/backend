package uc_upload

import (
	"github.com/tus/tusd/pkg/filestore"
	tusd "github.com/tus/tusd/pkg/handler"
	"log"
	"net/http"
	"os"
	"www.github.com/Maevlava/Matatani/backend/internal/middleware"
)

const BASE_PATH = "/upload/"

func TusHandler() (*tusd.Handler, error) {

	uploadDir := "./uploads"
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return nil, err
	}

	// later uses s3store or gcsstore
	store := filestore.FileStore{
		Path: uploadDir,
	}

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
