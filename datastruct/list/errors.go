package list

import "errors"

var (
	errEmptyListAccess = errors.New("error: accessed to empty list")
	errOutOfBounds     = errors.New("error: index out of bounds")
	errNegativeIndex   = errors.New("error: index must be positive")
)
