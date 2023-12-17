package batch

const (
	// Event constants
	AddIngredientEvent = "addIngredient"
	TakeSampleEvent    = "takeSample"
	StartBatchEvent    = "startBatch"
	EndBatchEvent      = "endBatch"

	// Status constants
	CreatedStatus = "created"
	RefingStatus  = "refining"
	EndedStatus   = "ended"

	// Ingredient constants
	CacaoNibs   = "cacao nibs"
	CacaoButter = "cacao butter"
	Sugar       = "sugar"
	MilkPowder  = "milk powder"
	Other       = "other"
)
