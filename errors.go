package gann

import "github.com/pkg/errors"

var (
	ErrDimensionMismatch         = errors.New("dimension mismatch")
	ErrInvalidIndex              = errors.New("invalid index")
	ErrItemNotFoundOnGivenItemID = errors.New("item not found for give item id")
)
