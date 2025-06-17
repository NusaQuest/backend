package main

import (
	"log"

	"github.com/NusaQuest/backend.git/config"
	"github.com/NusaQuest/backend.git/middleware"
	"github.com/NusaQuest/backend.git/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	app := fiber.New()

	config.ConnectDatabase()

	defer config.DisconnectDatabase()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	app.Use(middleware.CheckDBConnection)

	router.SetUp(app)

	log.Fatal(app.Listen(":8080"))
	
}