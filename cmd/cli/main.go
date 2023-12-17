package main

import (
	"fmt"

	"github.com/onasunnymorning/chocolatelab/pkg/batch"
)

func main() {
	// Create a Batch
	fmt.Println("Enter a batch name:")
	var name string
	fmt.Scanln(&name)
	b := batch.NewBatch(name)
	b.ID = 1 // Set the ID since we don't have a DB

	// Add some Cacao Butter
	fmt.Println("Amount of Cacao Butter to add:")
	var quantity int
	fmt.Scanln(&quantity)
	i, err := batch.NewIngredient("Cacao butter", batch.CacaoButter)
	if err != nil {
		panic(err)
	}
	e, err := batch.NewBatchEvent(b.ID, batch.AddIngredientEvent)
	if err != nil {
		panic(err)
	}
	e.Ingredient = *i
	e.Quantity = quantity
	b.AddEvent(e)

	// Add some nibs
	fmt.Println("Amount of Cacao Nibs to add:")
	fmt.Scanln(&quantity)
	i, err = batch.NewIngredient("Cacao nibs", batch.CacaoNibs)
	if err != nil {
		panic(err)
	}
	e, err = batch.NewBatchEvent(b.ID, batch.AddIngredientEvent)
	if err != nil {
		panic(err)
	}
	e.Ingredient = *i
	e.Quantity = quantity
	b.AddEvent(e)

	// Add some Sugar
	fmt.Println("Amount of Sugar to add:")
	fmt.Scanln(&quantity)
	i, err = batch.NewIngredient("Sugar", batch.Sugar)
	if err != nil {
		panic(err)
	}
	e, err = batch.NewBatchEvent(b.ID, batch.AddIngredientEvent)
	if err != nil {
		panic(err)
	}
	e.Ingredient = *i
	e.Quantity = quantity
	b.AddEvent(e)

	// Add some Milk Powder
	fmt.Println("Amount of Milk Powder to add:")
	fmt.Scanln(&quantity)
	i, err = batch.NewIngredient("Milk Powder", batch.MilkPowder)
	if err != nil {
		panic(err)
	}
	e, err = batch.NewBatchEvent(b.ID, batch.AddIngredientEvent)
	if err != nil {
		panic(err)
	}
	e.Ingredient = *i
	e.Quantity = quantity
	b.AddEvent(e)

	// Add some Other
	fmt.Println("Amount of Other to add:")
	fmt.Scanln(&quantity)
	i, err = batch.NewIngredient("Other", batch.Other)
	if err != nil {
		panic(err)
	}
	e, err = batch.NewBatchEvent(b.ID, batch.AddIngredientEvent)
	if err != nil {
		panic(err)
	}
	e.Ingredient = *i
	e.Quantity = quantity
	b.AddEvent(e)

	b.ParseEvents()

	for i := 0; i < len(b.Name)+10; i++ {
		fmt.Print("#")
	}
	fmt.Println()
	fmt.Printf("# Name: %s #\n", b.Name)
	for i := 0; i < len(b.Name)+10; i++ {
		fmt.Print("#")
	}
	fmt.Println()

	fmt.Printf("OutputWeight: %d\n", b.OutputWeight)
	fmt.Printf("CacaoPercentage: %f\n", b.CacaoPercentage)
	fmt.Printf("CacaoButterPercentage: %f\n", b.CacaoButterPercentage)
	fmt.Printf("Recipe: %+v\n", b.Recipe)

}
