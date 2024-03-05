package transport

import (
	"log"
	"net/http"
)

type httpServer struct {
	addr string
	mux  *http.ServeMux
}

func NewHttpServer(addr string) Server {
	return &httpServer{
		addr: addr,
		mux:  http.NewServeMux(),
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
