package cmd

import (
	"github.com/alidevjimmy/templaterenderer/internal/transport"
	"log"
)

func Serve() {
	server := transport.NewHttpServer(":8000")
	if err := server.Start(); err != nil {
		log.Fatalf("error: %v", err)
	}
}
