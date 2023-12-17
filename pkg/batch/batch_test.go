package batch

import "testing"

func TestNewBatch(t *testing.T) {
	name := "Hello\nWorld"
	expectedName := "Hello World"

	batch := NewBatch(name)

	if batch.Name != expectedName {
		t.Errorf("NewBatch(%q) = %q; expected %q", name, batch.Name, expectedName)
	}
}

func TestAddEvent(t *testing.T) {
	b := &Batch{}
	event := &BatchEvent{Type: AddIngredientEvent}

	b.AddEvent(event)

	if len(b.Events) != 1 {
		t.Errorf("AddEvent(%v) failed to add event to batch", event)
	}

	if b.Status != RefingStatus {
		t.Errorf("AddEvent(%v) failed to set correct status for addIngredientEvent", event)
	}

	event = &BatchEvent{Type: EndBatchEvent}

	b.AddEvent(event)

	if b.Status != EndedStatus {
		t.Errorf("AddEvent(%v) failed to set correct status for endBatchEvent", event)
	}
}

func TestParseEvents(t *testing.T) {
	b := &Batch{
		Events: []BatchEvent{
			{Type: AddIngredientEvent, Quantity: 1000, Ingredient: Ingredient{Type: CacaoNibs}},
			{Type: AddIngredientEvent, Quantity: 100, Ingredient: Ingredient{Type: CacaoButter}},
			{Type: AddIngredientEvent, Quantity: 500, Ingredient: Ingredient{Type: Sugar}},
			{Type: AddIngredientEvent, Quantity: 50, Ingredient: Ingredient{Type: MilkPowder}},
			{Type: AddIngredientEvent, Quantity: 50, Ingredient: Ingredient{Type: Other}},
		},
	}

	b.ParseEvents()

	expectedOutputWeight := 1700
	expectedCacaoWeight := 1000
	expectedCacaoButterWeight := 100
	expectedSugarWeight := 500
	expectedMilkWeight := 50
	expectedOtherWeight := 50
	expectedCacaoPercentage := (float64(expectedCacaoWeight) + float64(expectedCacaoButterWeight)) / float64(expectedOutputWeight) * 100
	expectedCacaoButterPercentage := float64(expectedCacaoButterWeight) / float64(expectedOutputWeight) * 100

	if b.OutputWeight != int(expectedOutputWeight) {
		t.Errorf("ParseEvents() failed, expected OutputWeight %d, got %d", expectedOutputWeight, b.OutputWeight)
	}

	if b.CacaoPercentage != expectedCacaoPercentage {
		t.Errorf("ParseEvents() failed, expected CacaoPercentage %f, got %f", expectedCacaoPercentage, b.CacaoPercentage)
	}

	if b.CacaoButterPercentage != expectedCacaoButterPercentage {
		t.Errorf("ParseEvents() failed, expected CacaoButterPercentage %f, got %f", expectedCacaoButterPercentage, b.CacaoButterPercentage)
	}

	expectedRecipe := Recipe{
		CacaoNibs:   int(expectedCacaoWeight),
		CacaoButter: int(expectedCacaoButterWeight),
		Sugar:       int(expectedSugarWeight),
		MilkPowder:  int(expectedMilkWeight),
		Other:       int(expectedOtherWeight),
	}

	if b.Recipe != expectedRecipe {
		t.Errorf("ParseEvents() failed, expected Recipe %+v, got %+v", expectedRecipe, b.Recipe)
	}
}
