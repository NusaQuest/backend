package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// @notice Represents an NFT ticket or asset linked to a concert event.
// @dev Used to store and retrieve NFT metadata from MongoDB.
type NFT struct {
	// @notice MongoDB Object ID (automatically generated)
	Id primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	// @notice Smart contract ID or on-chain identifier of the NFT
	ScId int64 `json:"scid" validate:"required"`

	// @notice Display name of the NFT
	Name string `json:"name" validate:"required"`

	// @notice Image URL or IPFS hash representing the NFT
	Image string `json:"image" validate:"required"`

	// @notice Description of the NFT content or purpose
	Description string `json:"description" validate:"required"`

	// @notice Name of the event this NFT is associated with
	ConcertName string `json:"concertname" validate:"required"`

	// @notice Date and time of the event in ISO 8601 format
	DateTime string `json:"datetime" validate:"required"`

	// @notice Extra benefits provided by the NFT (e.g., drinks, merch, VIP access)
	Utility string `json:"utility" validate:"required"`

	// @notice Seating zone or area associated with this ticket (e.g., Festival A, VIP Front Row)
	SeatZone string `json:"seatzone" validate:"required"`

	// @notice Price of the NFT in smallest unit of currency (e.g., wei)
	Price int64 `json:"price" validate:"required,min=1"`

	// @notice Total available stock of the NFT
	Stock int64 `json:"stock" validate:"required"`

	// @notice Number of NFTs that have been purchased
	Purchased int64 `json:"purchased"`
}
