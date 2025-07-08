// cmd/main.go
package main

import (
	"ELEVATE_INVIX_BE/configs"
	"ELEVATE_INVIX_BE/middlewares"
	"ELEVATE_INVIX_BE/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	configs.ConnectDB()

	app := fiber.New()

	// middleware
	app.Use(middlewares.Recovery())
	app.Use(middlewares.Logger())
	app.Use(middlewares.CORS())
	app.Use(middlewares.Helmet())
	app.Use(middlewares.RateLimiter())

	// Root route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(".... INVIX ....")
	})

	api := app.Group("/api")
	routes.AuthRoutes(api)	
	routes.UserRouter(api)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000" 
	}
	log.Fatal(app.Listen(":" + port))
}
