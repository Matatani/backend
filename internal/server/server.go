package server

import (
	"log"
	"net/http"
	"www.github.com/Maevlava/Matatani/backend/internal/config"
	"www.github.com/Maevlava/Matatani/backend/internal/predictor_service"
	"www.github.com/Maevlava/Matatani/backend/internal/uc_heading"
	"www.github.com/Maevlava/Matatani/backend/internal/uc_upload"
)

type MatataniServer struct {
	cfg             *config.APIConfig
	predictorClient predictor_service.PredictorClient
}

func NewMatataniServer(cfg *config.APIConfig, predictorClient predictor_service.PredictorClient) *MatataniServer {
	return &MatataniServer{
		cfg:             cfg,
		predictorClient: predictorClient,
	}
}
func (s *MatataniServer) NewHTTPRouter() http.Handler {
	mux := http.NewServeMux()

	// Heading features
	headingHandler := uc_heading.NewHandler()
	uc_heading.RegisterHeadingRoutes(mux, headingHandler)

	// Upload features
	tusHandler, err := uc_upload.TusHandler(s.cfg, s.predictorClient)
	uc_upload.RegisterUploadRoutes(mux, tusHandler)

	if err != nil {
		log.Fatalf("Unable to create tus handler: %v", err)
	}
	return mux
}
