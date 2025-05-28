package main

import (
	"context"
	"go-hexagonal-template/internal/adapters/http/handler"
	"go-hexagonal-template/internal/core/implementation/repository"
	"go-hexagonal-template/internal/core/implementation/services"
	"go-hexagonal-template/internal/core/ports"
	"go-hexagonal-template/internal/infrastructure/config"
	"go-hexagonal-template/internal/infrastructure/driven/envs"
	"go-hexagonal-template/internal/infrastructure/driven/mongodb"
	initserver "go-hexagonal-template/internal/server"
)

func main() {

	// Initialize the application
	var cfg config.AppConfig
	ctx := context.Background()
	ctx = envs.WithEnvs(ctx, &cfg)

	// Initialize database connection
	mongo := mongodb.NewMongoConnection(ctx, cfg.MongoUrl, cfg.MongoDB, cfg.AppName)
	defer mongo.DisconnectMongoDB(ctx)

	// initialize repositories
	usersRepo := repository.NewUsersRepository(mongo.GetDatabase())
	// initialize services
	usersService := services.NewUserService(usersRepo)
	//initialize hub service
	hub := ports.NewServiceHub(usersService)
	//Initialize handlers
	userHandler := handler.NewUsersHandler(hub)

	// Initialize the server
	server := initserver.NewServer(&initserver.Handlers{UsersHandler: userHandler})
	server.Start(cfg.Port)
}
