package controllers

import (
	"github.com/NusaQuest/backend.git/constants"
	"github.com/NusaQuest/backend.git/controllers/helper"
	"github.com/NusaQuest/backend.git/models"
	"github.com/NusaQuest/backend.git/output"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AddNFT handles POST /api/nfts
// @notice Stores a new NFT entry in the database.
// @param c Fiber context containing the NFT data in the request body.
// @return JSON response with insert result or error.
func AddNFT(c *fiber.Ctx) error {

	var nft models.NFT

	err := c.BodyParser(&nft)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, err.Error())
	}

	res, err := helper.InsertData(string(constants.NFTs), &nft)
	if err != nil {
		return output.GetError(c, fiber.StatusInternalServerError, err.Error())
	}

	return output.GetSuccess(c, string(constants.SuccessCreateMessage), fiber.Map{
		"result": res.InsertedID,
	})

}

// PurchaseNFT handles PATCH /api/nfts/:id
// @notice Increments the purchased count and decrements the stock of the specified NFT.
// @param c Fiber context containing the NFT ID as URL parameter.
// @return JSON response with update result or error.
func PurchaseNFT(c *fiber.Ctx) error {

	id := c.Params("id")
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, err.Error())
	}

	var nfts []models.NFT
	filter := bson.M{"_id": objId}

	_, err = helper.RetrieveData(filter, string(constants.NFTs), &nfts)
	if err != nil {
		return output.GetError(c, fiber.StatusInternalServerError, err.Error())
	}

	nfts[0].Purchased += 1
	nfts[0].Stock -= 1

	_, err = helper.UpdateData(string(constants.NFTs), "_id", objId, &nfts[0])
	if err != nil {
		return output.GetError(c, fiber.StatusInternalServerError, err.Error())
	}

	return output.GetSuccess(c, string(constants.SuccessUpdateMessage), fiber.Map{})

}

// GetNFTs handles GET /api/nfts
// @notice Retrieves all NFT entries from the database.
// @param c Fiber context for the request.
// @return JSON response with list of NFTs or error.
func GetNFTs(c *fiber.Ctx) error {

	var nfts []models.NFT

	_, err := helper.RetrieveData(bson.M{}, string(constants.NFTs), &nfts)
	if err != nil {
		return output.GetError(c, fiber.StatusInternalServerError, err.Error())
	}

	return output.GetSuccess(c, string(constants.SuccessGetMessage), fiber.Map{
		"nfts": nfts,
	})

}
