package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)


func CORS() fiber.Handler{
	cors_settings := cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:   "Origin, Content-Type, Accept, Authorization",
	})
	return cors_settings
}
