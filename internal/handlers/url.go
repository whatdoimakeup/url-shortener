package handlers

import "github.com/gofiber/fiber/v2"

func Greet(c *fiber.Ctx) error {
	return c.SendString("Hello!")
}

func ShortenUrl(c * fiber.Ctx) error {
	return c.Status(418).SendString("Shortener In development")
}

func RedirectToOriginal(c * fiber.Ctx) error {
	return c.Status(418).SendString("In development")
}