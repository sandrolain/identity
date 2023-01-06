package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sandrolain/identity/src-svc/waweb/clientgrpc"
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

	res, err := client.Login(context.Background(), &clientgrpc.LoginRequest{
		Email:    "sandrolain@outlook.com",
		Password: "test123456",
	})
	fmt.Printf("res: %v\n", res)

	app := fiber.New()

	app.Static("/", "./assets")

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello")
	})

	app.Listen(":3000")
}
