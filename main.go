package main

import (
	db "fiber-go/config"
	"fmt"

	routes "fiber-go/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Go sales api is running...")
	db.Connect()

	app := fiber.New()
	routes.Setup(app)

	app.Listen(":30001")
}
