package batch

import "gorm.io/gorm"

// BatchEvent is a struct that represents an event that occurs during the refinement of a batch of chocolate
type BatchEvent struct {
	gorm.Model
	BatchID    uint
	Type       string
	Notes      string
	Ingredient Ingredient
	Quantity   int // in grams
}

// NewBatchEvent is a constructor for BatchEvent
func NewBatchEvent(batchID uint, eventType string) (*BatchEvent, error) {
	if !contains(CleanInputString(eventType), validEventTypes) {
		return nil, ErrInvalidEventType
	}
	return &BatchEvent{
		BatchID: batchID,
		Type:    CleanInputString(eventType),
	}, nil
}
