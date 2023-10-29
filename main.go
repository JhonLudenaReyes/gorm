// @/main.go
package main

import (
	"log"

	"github.com/JhonLudenaReyes/gorm/config"
	"github.com/JhonLudenaReyes/gorm/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config.Connect()

	app.Get("/dogs", handlers.GetDogs)
	app.Get("/dogs/:id", handlers.GetDog)
	app.Post("/dogs", handlers.AddDog)
	app.Put("/dogs/:id", handlers.UpdateDog)
	app.Delete("/dogs/:id", handlers.RemoveDog)

	log.Fatal(app.Listen(":3000"))
}
