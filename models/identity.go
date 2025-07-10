package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// @notice Represents a successful KTP registration linked to a wallet.
// @dev Stores the hash of KTP data for verification purposes, without exposing personal information.
type Identity struct {
	// @notice MongoDB Object ID (automatically generated)
	Id primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	// @notice Wallet address that completed the KTP registration
	Wallet string `json:"wallet" validate:"required"`

	// @notice Hash of the user's KTP data (OCR result)
	Hash string `json:"hash" validate:"required"`

	// @notice Unix timestamp when the KTP registration was completed
	RegisteredAt int64 `json:"registeredat" validate:"required"`
}