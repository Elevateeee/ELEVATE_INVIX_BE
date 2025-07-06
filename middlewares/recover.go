package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)


func Recovery() fiber.Handler {
	return recover.New()
}
