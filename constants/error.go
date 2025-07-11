package constants

// Error defines a type for custom error messages used throughout the application.
// @notice Helps standardize and centralize error handling.
type Error string

const (
	// @notice Error shown when the MongoDB client is not properly initialized or connected.
	// @dev This typically indicates that `ConnectDatabase()` has not been called or failed.
	DatabaseNotInitialized Error = "MongoDB hasn’t been initialized or connected properly. Please check the database setup."

	// @notice Error returned when the application fails to ping the MongoDB server.
	// @dev Can be caused by network issues, incorrect URI, or MongoDB server not running.
	ErrorWhilePingToDatabase Error = "Unable to reach the MongoDB database. Please ensure it is running and accessible."

	// @notice Error returned when no identity is found for a given wallet address.
	// @dev Ensure the wallet has previously registered via the KTP registration process.
	IdentityNotFound Error = "No identity found for this wallet. Please register your KTP to continue."

	// @notice Error returned when a proposal fails AI validation (e.g. not a beach cleanup or invalid location).
	// @dev The proposal name/description/location did not match criteria for a valid beach cleanup initiative.
	ProposalValidationFailed Error = "Proposal input is not valid. Make sure the location exists and the activity clearly describes a beach cleanup."
)
