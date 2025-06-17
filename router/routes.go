package router

import (
	"github.com/NusaQuest/backend.git/controller"
	"github.com/gofiber/fiber/v2"
)

func SetUp(app *fiber.App) {

	app.Post("/api/auth/wallet", controller.ConnectWallet)

	app.Delete("/api/auth/wallet/:wallet", controller.DisconnectWallet)

	app.Post("/api/transactions", controller.AddTransaction)

	app.Get("/api/transactions/:wallet", controller.GetTransactions)

	app.Post("/api/destinations", controller.AddDestination)

	app.Get("/api/destinations", controller.GetDestinations)

}