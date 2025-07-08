package routes

import (
	"ELEVATE_INVIX_BE/controllers/admincontrollers"

	"github.com/gofiber/fiber/v2"
)

func AdminRouter(app fiber.Router) {
	adminRouter := app.Group("/admin")

	adminRouter.Get("/list", admincontrollers.ListAdmins)
	adminRouter.Post("/add", admincontrollers.AddAdmin)
}