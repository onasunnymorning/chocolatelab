package main

import (
	"fmt"

	"github.com/onasunnymorning/chocolatelab/pkg/batch"
)

func main() {
	b := batch.NewBatch("Test Batch")
	b.ID = 1

	// Add some Cacao Butter
	i, err := batch.NewIngredient("Cacao butter", batch.CacaoButter)
	if err != nil {
		panic(err)
	}
	e, err := batch.NewBatchEvent(b.ID, batch.AddIngredientEvent)
	if err != nil {
		panic(err)
	}
	e.Ingredient = *i
	e.Quantity = 100
	b.AddEvent(e)

	// Add some nibs
	i, err = batch.NewIngredient("Cacao nibs", batch.CacaoNibs)
	if err != nil {
		panic(err)
	}
	e, err = batch.NewBatchEvent(b.ID, batch.AddIngredientEvent)
	if err != nil {
		panic(err)
	}
	e.Ingredient = *i
	e.Quantity = 1000
	b.AddEvent(e)

	// Add some Sugar
	i, err = batch.NewIngredient("Sugar", batch.Sugar)
	if err != nil {
		panic(err)
	}
	e, err = batch.NewBatchEvent(b.ID, batch.AddIngredientEvent)
	if err != nil {
		panic(err)
	}
	e.Ingredient = *i
	e.Quantity = 500
	b.AddEvent(e)

	// Add some Milk Powder
	i, err = batch.NewIngredient("Milk Powder", batch.MilkPowder)
	if err != nil {
		panic(err)
	}
	e, err = batch.NewBatchEvent(b.ID, batch.AddIngredientEvent)
	if err != nil {
		panic(err)
	}
	e.Ingredient = *i
	e.Quantity = 50
	b.AddEvent(e)

	// Add some Other
	i, err = batch.NewIngredient("Other", batch.Other)
	if err != nil {
		panic(err)
	}
	e, err = batch.NewBatchEvent(b.ID, batch.AddIngredientEvent)
	if err != nil {
		panic(err)
	}
	e.Ingredient = *i
	e.Quantity = 50
	b.AddEvent(e)

	b.ParseEvents()

	fmt.Printf("OutputWeight: %d\n", b.OutputWeight)
	fmt.Printf("CacaoPercentage: %f\n", b.CacaoPercentage)
	fmt.Printf("CacaoButterPercentage: %f\n", b.CacaoButterPercentage)
	fmt.Printf("Recipe: %+v\n", b.Recipe)

}
