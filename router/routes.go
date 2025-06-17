package router

import (
	"github.com/NusaQuest/backend.git/controller"
	"github.com/gofiber/fiber/v2"
)

func SetUp(app *fiber.App) {

	app.Post("/api/auth/requests", controller.ConnectWallet)

	app.Post("/api/transactions/add", controller.AddTransaction)

	app.Get("/api/transactions/:wallet", controller.GetTransactions)

	app.Post("/api/destinations/add", controller.AddDestination)

	app.Get("/api/destinations", controller.GetDestinations)

}