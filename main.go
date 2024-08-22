package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// defino la structura de jason
type Usuario struct {
	nombre   string
	apellido string
}

/*
	type person struct {
		name string
		age  int
	}
*/
/*
func handleUser(c *fiber.Ctx) error {

	user := Usuario{
		nombre:   "Darwin",
		apellido: "Uzcategui",
	}

	return c.Status(fiber.StatusOK).JSON(user)

}
*/
func handleCreateUser(c *fiber.Ctx) error {

	user := Usuario{}
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	fmt.Println(user)
	return c.Status(fiber.StatusOK).JSON(user)

}

func main() {
	// fmt.Println("hola darwin")
	app := fiber.New()
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hola Darwin***!")
	})
	app.Get("/user", handleUser)
	//prueba de supernova
	
	//app.Post("/user", handleCreateUser)
	app.Listen(":3099")
}

func handleUser(c *fiber.Ctx) error {

	user := Usuario{
		nombre:   "Darwin",
		apellido: "Uzcategui",
	}

	return c.Status(fiber.StatusOK).JSON(user)

}
