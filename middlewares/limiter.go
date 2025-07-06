package middlewares

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)


func RateLimiter() fiber.Handler{
	limiterSettings := limiter.New(limiter.Config{
		Max:        100, 
		Expiration: 2 * time.Minute,
	})

	return limiterSettings
}
