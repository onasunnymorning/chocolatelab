package batch

import "time"

// Origin is a struct that represents the origin of an ingredient
type Origin struct {
	ID          uint `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Ingredients []Ingredient
}

// NewOrigin is a constructor for Origin
func NewOrigin(name string) *Origin {
	return &Origin{
		Name: CleanInputString(name),
	}
}
