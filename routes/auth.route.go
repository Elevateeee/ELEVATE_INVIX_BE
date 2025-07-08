package routes

import (
	"ELEVATE_INVIX_BE/controllers/usercontrollers"
	"ELEVATE_INVIX_BE/middlewares"

	"github.com/gofiber/fiber/v2"
)


func AuthRoutes(app fiber.Router) {
	auth := app.Group("/auth")

	auth.Post("/register", usercontrollers.RegisterUser)
	auth.Post("/verify-email", usercontrollers.VerifyEmail)
	auth.Post("/resend-verification", usercontrollers.ResendVerification)
	auth.Post("/login", usercontrollers.Login)
	auth.Post("/logout",middlewares.UserProtect ,usercontrollers.Logout)
}


// note:
// http://127.0.0.1:3888/api/auth/register
// http://127.0.0.1:3888/api/auth/verify-email
// http://127.0.0.1:3888/api/auth/resend-verification
// http://127.0.0.1:3888/api/auth/login
// http://127.0.0.1:3888/api/auth/logout
