package middlewares

import (
	"github.com/Zain0205/cf-stunting-backend-go/internal/config"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func JWTProtected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(config.JWTSecret()),
	})
}
