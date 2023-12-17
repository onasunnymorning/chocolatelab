package batch

import "gorm.io/gorm"

// Ingredient is a struct that represents an ingredient in a batch
type Ingredient struct {
	gorm.Model
	Name     string
	OriginID uint
	Type     string
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
