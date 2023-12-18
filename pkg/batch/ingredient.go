package batch

import "time"

// Ingredient is a struct that represents an ingredient in a batch
type Ingredient struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	OriginID  uint
	Type      string
}

// NewIngredient is a constructor for Ingredient
func NewIngredient(name string, t string) (*Ingredient, error) {
	if !contains(CleanInputString(t), validIngredients) {
		return nil, ErrInvalidIngredient
	}
	return &Ingredient{
		Name: CleanInputString(name),
		Type: t,
	}, nil
}
