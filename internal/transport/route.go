package transport

import (
	"github.com/alidevjimmy/templaterenderer/internal/handler"
)

func (h *httpServer) routes() {
	h.mux.HandleFunc("GET /", handler.Home)
	h.mux.HandleFunc("GET /profile/{username}", h.userHandler.Profile)
	h.mux.HandleFunc("POST /profile", h.userHandler.Join)
}
