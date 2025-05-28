package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-hexagonal-template/internal/core/ports"
)

func UsersRoute(router fiber.Router, handler ports.IUserHandler) {
	router.Get("", handler.GetUsers)
	router.Post("/insert", handler.InsertUser)
}
