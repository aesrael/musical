package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func httpServer() *fiber.App {
	app := fiber.New()

	api := app.Group("/api", auth)

	app.Use(logger.New())
	app.Use(requestid.New())

	api.Post("/job", enqueueJob)
	api.Get("/backup", backupDb)

	return app
}
