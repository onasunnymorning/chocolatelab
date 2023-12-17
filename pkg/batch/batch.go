package batch

import "gorm.io/gorm"

// Batch is a struct that represents a batch of chocolate as it is refined
type Batch struct {
	gorm.Model
	Name                  string
	Notes                 string
	Events                []BatchEvent
	Status                string
	OutputWeight          int // in grams
	CacaoPercentage       float64
	CacaoButterPercentage float64
	Recipe                Recipe
}

// NewBatch is a constructor for Batch
func NewBatch(name string) *Batch {
	return &Batch{
		Name:   CleanInputString(name),
		Status: CreatedStatus,
	}
}

// AddEvent adds a new event to a batch. It will update the status of the batch if necessary. Adding the first ingredient will cause the status to become refining. Ending the batch will cause the status to become ended.
func (b *Batch) AddEvent(event *BatchEvent) {
	b.Events = append(b.Events, *event)
	// Modify the status only for interesting events
	switch event.Type {
	case AddIngredientEvent:
		b.Status = RefingStatus
	case EndBatchEvent:
		b.Status = EndedStatus
	}
}

// ParseEvents parses the events in a batch and updates the batch's outputWeight, cacao percentage, cacao butter percentage and recipe
func (b *Batch) ParseEvents() {
	var outputWeight, cacaoWeight, cacaoButterWeight, sugarWeight, milkWeight, otherWeight float64
	// loop over events and set counters
	for _, event := range b.Events {
		if event.Type == AddIngredientEvent {
			outputWeight += float64(event.Quantity)
			switch event.Ingredient.Type {
			case CacaoNibs:
				cacaoWeight += float64(event.Quantity)
			case CacaoButter:
				cacaoButterWeight += float64(event.Quantity)
			case Sugar:
				sugarWeight += float64(event.Quantity)
			case MilkPowder:
				milkWeight += float64(event.Quantity)
			default:
				otherWeight += float64(event.Quantity)
			}
		}
	}
	// calculate percentages
	b.OutputWeight = int(outputWeight)
	b.CacaoPercentage = (cacaoWeight + cacaoButterWeight) / outputWeight * 100
	b.CacaoButterPercentage = cacaoButterWeight / outputWeight * 100

	// create the recipe
	b.Recipe = Recipe{
		CacaoNibs:   int(cacaoWeight),
		CacaoButter: int(cacaoButterWeight),
		Sugar:       int(sugarWeight),
		MilkPowder:  int(milkWeight),
		Other:       int(otherWeight),
	}
}
