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
	"github.com/joho/godotenv"
)

// @notice Entry point of the NusaQuest backend application.
// @dev Sets up Fiber server, connects to MongoDB, registers routes, middleware, and starts the HTTP server.
func main() {
	// @notice Initialize HTML template engine
	engine := html.New("./views", ".html")

	// @notice Create new Fiber app with custom template engine
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// @notice Load environment variables from .env file into runtime
	// @dev Used to configure secrets, database URL, and other env-specific values
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Cannot load .env! %s", err)
	}

	// @notice Connect to MongoDB using configuration
	config.ConnectDatabase()

	// @dev Ensure DB connection is closed on shutdown
	defer config.DisconnectDatabase()

	// @notice Enable CORS for all origins and headers (not secure for production)
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	// @notice Middleware to check database connection on each request
	app.Use(middlewares.CheckDBConnection)

	// @notice Set up application routes
	router.SetUp(app)

	// @notice Custom 404 handler for undefined routes
	app.Use(handlers.NotFound)

	// @notice Start the Fiber server on port 8080
	log.Fatal(app.Listen(":8080"))
}
