package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// @notice Represents a blockchain-related transaction log associated with a user wallet.
// @dev Used to store and retrieve transaction history from MongoDB.
type Transaction struct {
	// @notice MongoDB Object ID (automatically generated)
	Id primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	// @notice Wallet address that performed the transaction
	Wallet string `json:"wallet" validate:"required"`

	// @notice Short title or summary of the transaction
	Title string `json:"title" validate:"required"`

	// @notice Detailed description of what the transaction is for
	Detail string `json:"detail" validate:"required"`

	// @notice Hash of the transaction recorded on blockchain
	TxHash string `json:"txhash" validate:"required"`

	// @notice Unix timestamp when the transaction occurred
	TxTimestamp uint64 `json:"txtimestamp" validate:"required"`
}
