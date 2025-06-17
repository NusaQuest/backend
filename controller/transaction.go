package controller

import (
	"github.com/NusaQuest/backend.git/output"
	"github.com/gofiber/fiber/v2"
)

func AddTransaction(c *fiber.Ctx) error {

	return output.GetSuccess(c, "", fiber.Map{
		"transaction": "",
	})

}

func GetTransactions(c *fiber.Ctx) error {

	return output.GetSuccess(c, "", fiber.Map{
		"transactions": "",
	})

}