package handlers

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/sandrolain/identity/src-svc/waweb/clientgrpc"
)

var grpcCLient clientgrpc.ClientServiceClient

func InitHandlers(app *fiber.App, grpc clientgrpc.ClientServiceClient) {
	grpcCLient = grpc
	app.Post("/api/login", Login)
	app.Post("/api/loginConfirm", LoginConfirm)
}

func getBearerToken(c *fiber.Ctx) (token string, err error) {
	auth := c.Get("Authorization")
	parts := strings.Split(auth, " ")
	if len(parts) > 1 {
		token = parts[1]
	}
	if token == "" {
		err = fmt.Errorf("empty authorization token")
	}
	return
}

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type LoginConfirmRequest struct {
	TotpCode string `json:"totpCode" form:"totpCode"`
}

func Login(c *fiber.Ctx) (err error) {
	var req LoginRequest
	if err = c.BodyParser(&req); err != nil {
		return
	}
	res, err := grpcCLient.Login(c.UserContext(), &clientgrpc.LoginRequest{Email: req.Email, Password: req.Password})
	if err != nil {
		return err
	}
	return c.JSON(res)
}

func LoginConfirm(c *fiber.Ctx) (err error) {
	var req LoginConfirmRequest
	if err = c.BodyParser(&req); err != nil {
		return
	}
	totpToken, err := getBearerToken(c)
	if err != nil {
		return
	}
	greq := &clientgrpc.LoginConfirmRequest{TotpToken: totpToken, TotpCode: req.TotpCode}
	fmt.Printf("greq: %+v\n", *greq)
	res, err := grpcCLient.LoginConfirm(c.UserContext(), greq)
	if err != nil {
		return
	}
	return c.JSON(res)
}
