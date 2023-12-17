package batch

import "errors"

var (
	ErrInvalidEventType  = errors.New("invalid event type see validEventTypes for valid types")
	ErrInvalidStatus     = errors.New("invalid status see validStatuses for valid statuses")
	ErrInvalidIngredient = errors.New("invalid ingredient see validIngredients for valid ingredients")
)
