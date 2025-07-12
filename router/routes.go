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

	// @notice Delete an existing proposal by ID
	app.Delete("/api/proposals/:id", controllers.DeleteProposal)

	// @notice Get all proposals
	app.Get("/api/proposals", controllers.GetProposals)

	// @notice Add a new nft entry
	app.Post("/api/nfts", controllers.AddNFT)

	// @notice Purchase a specific NFT by ID (increments purchased count and decreases stock)
	app.Patch("/api/nfts/:id", controllers.PurchaseNFT)

	// @notice Get all proposals
	app.Get("/api/nfts", controllers.GetNFTs)

}
