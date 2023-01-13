package handlers

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/sandrolain/identity/src-svc/waweb/clientgrpc"
)

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

type LoginConfirmRequest struct {
	TotpCode string `json:"totpCode" form:"totpCode"`
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
	res, err := grpcCLient.LoginConfirm(c.UserContext(), greq)
	if err != nil {
		return
	}
	return c.JSON(res)
}

func BeginWebauthnRegister(c *fiber.Ctx) (err error) {
	sessionToken, err := getBearerToken(c)
	if err != nil {
		return
	}
	res, err := grpcCLient.BeginWebauthnRegister(c.UserContext(), &clientgrpc.BeginWebauthnRegisterRequest{
		SessionToken: sessionToken,
	})
	if err != nil {
		return err
	}

	var r interface{}
	err = json.Unmarshal([]byte(res.CredentialCreation), &r)
	if err != nil {
		return err
	}

	return c.JSON(r)
}

func FinishWebauthnRegister(c *fiber.Ctx) (err error) {
	sessionToken, err := getBearerToken(c)
	if err != nil {
		return
	}
	res, err := grpcCLient.FinishWebauthnRegister(c.UserContext(), &clientgrpc.FinishWebauthnRegisterRequest{
		SessionToken: sessionToken,
		Request:      c.Body(),
	})
	if err != nil {
		return err
	}
	return c.JSON(res)
}

type BeginWebauthnLoginRequest struct {
	Email string `json:"email" form:"email"`
}

type BeginWebauthnLoginResponse struct {
	WebauthnToken       string      `json:"webauthnToken"`
	CredentialAssertion interface{} `json:"credentialAssertion"`
}

func BeginWebauthnLogin(c *fiber.Ctx) (err error) {
	var req BeginWebauthnLoginRequest
	if err = c.BodyParser(&req); err != nil {
		return
	}
	res, err := grpcCLient.BeginWebauthnLogin(c.UserContext(), &clientgrpc.BeginWebauthnLoginRequest{
		Email: req.Email,
	})
	if err != nil {
		return err
	}

	var r interface{}
	err = json.Unmarshal([]byte(res.CredentialAssertion), &r)
	if err != nil {
		return err
	}

	return c.JSON(BeginWebauthnLoginResponse{
		WebauthnToken:       res.WebauthnToken,
		CredentialAssertion: r,
	})
}

func FinishWebauthnLogin(c *fiber.Ctx) (err error) {
	waToken, err := getBearerToken(c)
	if err != nil {
		return
	}
	res, err := grpcCLient.FinishWebauthnLogin(c.UserContext(), &clientgrpc.FinishWebauthnLoginRequest{
		WebauthnToken: waToken,
		Request:       c.Body(),
	})
	if err != nil {
		return err
	}
	return c.JSON(res)
}
