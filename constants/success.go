package constants

// Success defines a type for success messages returned by the application.
// @notice Used to provide standardized responses for successful operations.
type Success string

const (
	// @notice Message returned when data is successfully retrieved.
	// @dev Commonly used in GET API responses.
	SuccessGetMessage Success = "Data successfully retrieved."

	// @notice Message returned when data is successfully created.
	// @dev Commonly used in POST API responses.
	SuccessCreateMessage Success = "Data successfully created."

	// @notice Message returned when data is successfully deleted.
	// @dev Commonly used in DELETE API responses.
	SuccessDeleteMessage Success = "Data successfully deleted."

	// @notice Message returned when data is successfully updated.
	// @dev Commonly used in PUT/PATCH API responses.
	SuccessUpdateMessage Success = "Data successfully updated."
)
