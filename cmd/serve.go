package cmd

import (
	"github.com/alidevjimmy/templaterenderer/internal/handler"
	"github.com/alidevjimmy/templaterenderer/internal/repository"
	"github.com/alidevjimmy/templaterenderer/internal/service"
	"github.com/alidevjimmy/templaterenderer/internal/transport"
	"log"
)

func Serve() {
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	server := transport.NewHttpServer(":8000", userHandler)

	if err := server.Start(); err != nil {
		log.Fatalf("error: %v", err)
	}
}
