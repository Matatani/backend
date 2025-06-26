package uc_heading

import (
	"net/http"
	"www.github.com/Maevlava/Matatani/backend/internal/common"
	"www.github.com/Maevlava/Matatani/backend/internal/middleware"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) GetHeading(w http.ResponseWriter, r *http.Request) {
	data := HeadingResponse{
		Text: "Welcome to Matatani: The Main Event!",
	}
	common.RespondWithJSON(w, http.StatusOK, data)
}

func RegisterHeadingRoutes(router *http.ServeMux, h *Handler) {

	// Middleware : CORS
	getHeadingHandler := middleware.EnableCORS(http.HandlerFunc(h.GetHeading))

	router.HandleFunc("/heading", getHeadingHandler)
}
