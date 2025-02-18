package main

import "github.com/NATALI-LEV/URL-SHORTENER/api/routes"

func setupRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL)
}

func main() {

}
