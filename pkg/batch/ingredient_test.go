package batch

import "testing"

func TestNewIngredient_ValidIngredient(t *testing.T) {
	name := "Chocolate"
	ingredientType := CacaoNibs
	ingredient, err := NewIngredient(name, ingredientType)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if ingredient.Name != name {
		t.Errorf("Expected ingredient name to be %s, got %s", name, ingredient.Name)
	}

	if ingredient.Type != ingredientType {
		t.Errorf("Expected ingredient type to be %s, got %s", ingredientType, ingredient.Type)
	}
}

func TestNewIngredient_InvalidIngredient(t *testing.T) {
	name := "Chocolate"
	ingredientType := "Invalid"
	_, err := NewIngredient(name, ingredientType)
	if err != ErrInvalidIngredient {
		t.Errorf("Expected error %v, got %v", ErrInvalidIngredient, err)
	}
}
