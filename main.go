package main

import (
	"log"

	"github.com/NusaQuest/backend.git/config"
	"github.com/NusaQuest/backend.git/handlers"
	"github.com/NusaQuest/backend.git/middlewares"
	"github.com/NusaQuest/backend.git/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"
)

func main() {

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	config.ConnectDatabase()

	defer config.DisconnectDatabase()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	app.Use(middlewares.CheckDBConnection)

	router.SetUp(app)

	app.Use(handlers.NotFound)

	log.Fatal(app.Listen(":8080"))
	
}