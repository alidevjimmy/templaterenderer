package transport

import (
	"github.com/alidevjimmy/templaterenderer/internal/handler"
	"log"
	"net/http"
)

type httpServer struct {
	addr        string
	mux         *http.ServeMux
	userHandler *handler.UserHandler
}

func NewHttpServer(addr string, userHandler *handler.UserHandler) Server {
	return &httpServer{
		addr:        addr,
		mux:         http.NewServeMux(),
		userHandler: userHandler,
	}
}

func (h *httpServer) Start() error {
	h.routes()
	s := http.Server{
		Addr:    h.addr,
		Handler: h.mux,
	}
	log.Println("starting http server")
	return s.ListenAndServe()
}

func (h *httpServer) Shutdown() error {
	return nil
}
