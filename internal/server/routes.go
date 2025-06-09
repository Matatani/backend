package server

import (
	"net/http"
	"www.github.com/Maevlava/Matatani/backend/internal/heading"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	// Heading features
	headingHandler := heading.NewHandler()
	heading.RegisterHeadingRoutes(mux, headingHandler)

	return mux
}
