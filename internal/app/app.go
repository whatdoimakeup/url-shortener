package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/whatdoimakeup/url-shortener/internal/routes"
)

func Run() error {
	app := fiber.New(fiber.Config{
		Prefork: false,
	})

	app.Use(logger.New())
	app.Use(recover.New())
	routes.RegisterRoutes(app)

	return app.Listen(":8000") // Replace with Config
}