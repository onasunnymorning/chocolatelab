package batch

import "gorm.io/gorm"

// Origin is a struct that represents the origin of an ingredient
type Origin struct {
	gorm.Model
	Name        string
	Ingredients []Ingredient
}

// NewOrigin is a constructor for Origin
func NewOrigin(name string) *Origin {
	return &Origin{
		Name: CleanInputString(name),
	}
}
