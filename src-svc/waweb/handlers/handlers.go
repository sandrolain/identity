package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sandrolain/identity/src-svc/waweb/clientgrpc"
)

var grpcCLient clientgrpc.ClientServiceClient

func InitHandlers(app *fiber.App, grpc clientgrpc.ClientServiceClient) {
	grpcCLient = grpc
	app.Post("/api/login", Login)
	app.Post("/api/loginConfirm", LoginConfirm)
}
