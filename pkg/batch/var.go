package batch

var (
	validEventTypes = []string{
		AddIngredientEvent,
		TakeSampleEvent,
		StartBatchEvent,
		EndBatchEvent,
	}

	validIngredients = []string{
		CacaoNibs,
		CacaoButter,
		Sugar,
		MilkPowder,
		Other,
	}
)
