package main

import (
	"github.com/NATALI-LEV/URL-SHORTENER/api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL)
}

func main() {
	err := godotenv.Load()
	app := fiber.New()

}
