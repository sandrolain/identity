package main

import (
	"context"
	"fmt"
	"log"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/sandrolain/go-utilities/pkg/httputils"
	"github.com/sandrolain/identity/src-svc/waweb/clientgrpc"
	"github.com/sandrolain/identity/src-svc/waweb/handlers"
)

func main() {
	serverAddr := "localhost:1985"
	caFile := "../../localtest/cert/ca-cert.pem"
	serverHostOverride := "waweb.sandrolain.com"

	client, conn, err := clientgrpc.NewClient(serverAddr, caFile, serverHostOverride)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	app := fiber.New()

	app.Static("/", "./assets")

	handlers.InitHandlers(app, client)

	var totpUri string

	app.Get("/generateSessionToken", func(c *fiber.Ctx) error {
		res, err := client.Login(context.Background(), &clientgrpc.LoginRequest{
			Email:    "sandrolain@outlook.com",
			Password: "test123456",
		})
		if err != nil {
			return err
		}

		if res.TotpUri != "" {
			totpUri = res.TotpUri
		}

		url := fmt.Sprintf("https://www.sandrolain.com/.netlify/functions/totpCode?uri=%v", url.QueryEscape(totpUri))
		totp, err := httputils.Fetch(url).BodyString()
		if err != nil {
			return err
		}

		res2, err := client.LoginConfirm(context.Background(), &clientgrpc.LoginConfirmRequest{
			TotpToken: res.TotpToken,
			TotpCode:  totp,
		})
		if err != nil {
			return err
		}

		return c.JSON(res2)
	})

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello")
	})

	err = app.Listen(":3000")
	fmt.Print(err)
}
