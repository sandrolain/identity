package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
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

	err = app.Listen(":3000")
	fmt.Print(err)
}
