package constants

// Document defines the type for MongoDB collection names.
// @notice This type is used to ensure type safety when referring to collection names.
type Document string

const (
	// @notice The MongoDB collection name for storing proposal data.
	Proposals Document = "proposals"

	// @notice The MongoDB collection name for storing transaction data.
	Transactions Document = "transactions"

	// @notice The MongoDB collection name for storing identity data.
	Identities Document = "identities"
)
