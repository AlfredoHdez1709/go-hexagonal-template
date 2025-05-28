package handler

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"go-hexagonal-template/internal/core/domain"
	"go-hexagonal-template/internal/core/ports"
	errorsConst "go-hexagonal-template/internal/infrastructure/constants/errors"
)

type UsersHandler struct {
	ServiceHub ports.ServiceHub
}

func (u UsersHandler) InsertUser(c *fiber.Ctx) error {
	ctx := c.Context()
	var user domain.User
	if err := c.BodyParser(&user); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(errorsConst.NewHTTPDefaultExceptionResponse(fiber.StatusBadRequest, "Invalid request body"))
	}

	if err := u.ServiceHub.UserService.InsertUser(ctx, user); err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(errorsConst.NewHTTPDefaultExceptionResponse(fiber.StatusInternalServerError, err.Error()))
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(user)
}

func (u UsersHandler) GetUsers(c *fiber.Ctx) error {
	ctx := c.Context()
	users, err := u.ServiceHub.UserService.GetUsers(ctx)
	if err != nil {
		if errors.Is(err, errorsConst.ErrNotFound) {
			c.Status(fiber.StatusNoContent)
			return c.JSON(errorsConst.NewHTTPDefaultExceptionResponse(fiber.StatusNotFound, errorsConst.ErrNotFound.Error()))
		}
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(errorsConst.NewHTTPDefaultExceptionResponse(fiber.StatusInternalServerError, err.Error()))
	}

	return c.JSON(users)
}

func NewUsersHandler(hub ports.ServiceHub) ports.IUserHandler {
	return &UsersHandler{
		ServiceHub: hub,
	}
}
