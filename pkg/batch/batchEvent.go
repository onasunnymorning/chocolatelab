package batch

import "time"

// BatchEvent is a struct that represents an event that occurs during the refinement of a batch of chocolate
type BatchEvent struct {
	ID           uint `gorm:"primaryKey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	BatchID      uint
	Type         string
	Notes        string
	IngredientID uint
	Ingredient   Ingredient
	Quantity     int // in grams
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
