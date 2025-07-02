package router

import (
	"github.com/NusaQuest/backend.git/controllers"
	"github.com/NusaQuest/backend.git/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetUp(app *fiber.App) {

	app.Get("/", handlers.Welcome)

	app.Post("/api/transactions", controllers.AddTransaction)

	app.Get("/api/transactions/:wallet", controllers.GetWalletTransactions)

	app.Post("/api/proposals", controllers.AddProposal)

	app.Patch("/api/proposals/:id", controllers.UpdateProposal)

	app.Get("/api/proposals", controllers.GetProposals)

	app.Get("/api/proposals/:wallet", controllers.GetWalletProposals)

}
