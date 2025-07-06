package router

import (
	"github.com/NusaQuest/backend.git/controllers"
	"github.com/NusaQuest/backend.git/handlers"
	"github.com/gofiber/fiber/v2"
)

// @notice Set up all application routes for the NusaQuest backend.
// @dev Registers endpoints for transaction and proposal handling.
// @param app The main Fiber application instance.
func SetUp(app *fiber.App) {

	// @notice Root route to check if the server is running
	app.Get("/", handlers.Welcome)

	// @notice Add a new identity entry
	app.Post("/api/identities", controllers.RegisterIdentity)

	// @notice Get identity submitted by a specific wallet
	app.Get("/api/identities/:wallet", controllers.GetIdentity)

	// @notice Add a new transaction entry
	app.Post("/api/transactions", controllers.AddTransaction)

	// @notice Get all transactions by a wallet address
	app.Get("/api/transactions/:wallet", controllers.GetWalletTransactions)

	// @notice Add a new proposal entry
	app.Post("/api/proposals", controllers.AddProposal)

	// @notice Update an existing proposal by ID
	app.Patch("/api/proposals/:id", controllers.UpdateProposal)

	// @notice Get all proposals
	app.Get("/api/proposals", controllers.GetProposals)

	// @notice Get all proposals submitted by a specific wallet
	app.Get("/api/proposals/:wallet", controllers.GetWalletProposals)
}
