package routes

import (
	"ELEVATE_INVIX_BE/controllers/usercontrollers"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(app fiber.Router) {
	userRouter := app.Group("/user")

	userRouter.Get("/list", usercontrollers.ListUsers)
	userRouter.Post("/add", usercontrollers.AddUser)
}
