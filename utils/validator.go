package utils

import "github.com/go-playground/validator/v10"

// @notice Initializes and returns a new instance of the validator.
// @dev Uses go-playground/validator.v10 for struct validation.
// @return Pointer to a new validator.Validate instance.
func GetValidator() *validator.Validate {
	return validator.New()
}
