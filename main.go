package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("hola darwin")
	app := fiber.New()
	app.Listen(":3000")
}
