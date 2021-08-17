package main

import (
	"github.com/gofiber/fiber/v2"
)

func httpServer() *fiber.App {
	app := fiber.New()

	api := app.Group("/api", auth)

	api.Get("/job", enqueueJob)
	// api.Post("/backup")

	return app
}
