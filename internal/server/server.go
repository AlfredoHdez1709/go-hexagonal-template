package server

import (
	"fmt"
	"go-hexagonal-template/internal/adapters/http/routes"
	"go-hexagonal-template/internal/core/ports"
	"go-hexagonal-template/internal/infrastructure/driven/fiber"
	"go-hexagonal-template/internal/infrastructure/driven/logger"
)

type Server struct {
	Fiber    *fiber.FiberServer
	Handlers *Handlers
}

type Handlers struct {
	UsersHandler ports.IUserHandler
}

func NewServer(handlers *Handlers) *Server {
	fiberServer := fiber.NewFiberServer()
	return &Server{
		Fiber:    fiberServer,
		Handlers: handlers,
	}
}

func (server *Server) Start(port string) {
	api := server.Fiber.Server.Group("/api/v1")

	//example routes
	usersGroup := api.Group("/users")
	routes.UsersRoute(usersGroup, server.Handlers.UsersHandler)

	// Add more routes as needed

	err := server.Fiber.Server.Listen(fmt.Sprintf(":%s", port))
	if err != nil {
		logger.Logger.Fatal(err)
	}
}
