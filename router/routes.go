package router

import (
	"github.com/NusaQuest/backend.git/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetUp(app *fiber.App) {

	app.Post("/api/transactions", controllers.AddTransaction)

	app.Get("/api/transactions/:wallet", controllers.GetTransactions)

	app.Post("/api/destinations", controllers.AddDestination)

	app.Get("/api/destinations", controllers.GetDestinations)

}