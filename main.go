package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/samarth-srivastava/social-media/controllers"
	"github.com/samarth-srivastava/social-media/database"
	"github.com/samarth-srivastava/social-media/middleware"
)

func main() {
	app := fiber.New(fiber.Config{
        ErrorHandler: middleware.ErrorHandlerMiddlerware,
    })
	app.Use(cors.New());
	app.Use(middleware.LoggerMiddleware)
	database.Connect();
	fmt.Println("Connected to MongoDB");
	setupRoutes(app);
	port := os.Getenv("PORT")
	if port == ""{
		port = "5000"
	}
	app.Listen(":" + port);
}

func setupRoutes(app *fiber.App){
	app.Post("/api/register",controllers.Register);
	app.Post("/api/login",controllers.Login);

	api := app.Use(middleware.Middleware);

	//post endpoints
	api.Post("/api/posts",controllers)
}