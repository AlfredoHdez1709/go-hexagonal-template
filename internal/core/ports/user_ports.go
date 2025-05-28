package ports

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go-hexagonal-template/internal/core/domain"
)

type IUserHandler interface {
	GetUsers(c *fiber.Ctx) error
	InsertUser(c *fiber.Ctx) error
}

type IUserService interface {
	GetUsers(ctx context.Context) ([]domain.User, error)
	InsertUser(ctx context.Context, user domain.User) error
}

type IUserRepository interface {
	GetUsers(ctx context.Context) ([]domain.User, error)
	InsertUser(ctx context.Context, user domain.User) error
}
