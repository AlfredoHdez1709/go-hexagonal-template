package fiber

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	errorsConst "go-hexagonal-template/internal/infrastructure/constants/errors"
)

type FiberServer struct {
	Server *fiber.App
}

func NewFiberServer() *FiberServer {
	server := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}
			return ctx.Status(code).JSON(errorsConst.MessageError{
				Status:  e.Code,
				Message: err.Error(),
			})
		},
	})

	server.Use(cors.New())
	return &FiberServer{Server: server}
}

func (f *FiberServer) Start(port string) error {
	return f.Server.Listen(fmt.Sprintf(":%v", port))
}
