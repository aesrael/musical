package main

import (
	"musical/config"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func auth(c *fiber.Ctx) error {
	tokenStr := c.Get("bearerToken")
	if tokenStr == "" {
		c.SendStatus(http.StatusUnauthorized)
		return nil
	}

	_, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config["JWT_KEY"]), nil
	})

	if err != nil {
		c.Status(http.StatusUnauthorized).SendString(err.Error())
		return err
	}
	return c.Next()
}
