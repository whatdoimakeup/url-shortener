package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/whatdoimakeup/url-shortener/internal/handlers"
)

func RegisterRoutes(app *fiber.App) {
	app.Get("/", handlers.Greet)
	app.Post("/shorten", handlers.ShortenUrl)
	app.Get("/:short", handlers.RedirectToOriginal)
}