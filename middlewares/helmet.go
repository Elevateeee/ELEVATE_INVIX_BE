package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
)

func Helmet()  fiber.Handler {
	helmetSetting := helmet.New();
	return helmetSetting
}
