package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type SomeStruct struct {
	Name string
	Age  uint8
}

func main2() {
	app := fiber.New()

	app.Use(cors.New())

	/*
		app.Use(func(c *fiber.Ctx) error {
			if c.Is("json") {
				return c.Next()
			}
			return c.SendString("Only JSON allowed!")
		})
	*/
	app.Get("/pepe", func(c *fiber.Ctx) error {
		// Create data struct:
		data := SomeStruct{
			Name: "Grame",
			Age:  20,
		}
		return c.Status(fiber.StatusOK).JSON(data)
		//return c.JSON(data)
	})

	log.Fatal(app.Listen(":3000"))
}
