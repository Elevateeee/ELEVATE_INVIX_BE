package routes

import (
	"ELEVATE_INVIX_BE/controllers/authcontrollers"
	"ELEVATE_INVIX_BE/middlewares"

	"github.com/gofiber/fiber/v2"
)


func AuthRoutes(app fiber.Router) {
	auth := app.Group("/auth")

	auth.Post("/register", authcontrollers.RegisterUser)
	auth.Post("/verify-email", authcontrollers.VerifyEmail)
	auth.Post("/resend-verification", authcontrollers.ResendVerification)
	auth.Post("/login", authcontrollers.Login)
	auth.Post("/logout",middlewares.UserProtect ,authcontrollers.Logout)
}


// note:
// http://127.0.0.1:3888/api/auth/register
// http://127.0.0.1:3888/api/auth/verify-email
// http://127.0.0.1:3888/api/auth/resend-verification
// http://127.0.0.1:3888/api/auth/login
// http://127.0.0.1:3888/api/auth/logout
