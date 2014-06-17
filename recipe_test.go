package main

import (
	"testing"
	// "fmt"
)

// Test the Recipe struct and access methods
func TestRecipeStruct(t *testing.T) {
	var newRecipe Recipe
	newRecipe = Recipe{ 
		Title: "My new recipe", 
		Source: "Brinker Family", 
		Yield: "6 servings", 
		Description: "It's good",
		Images: []string{"image1.jpg", "image2.jpg"},
		IngredientLists: []IngredientSection{ 
			IngredientSection{
				Title: "Frosting", 
				Ingredients: []Ingredient{
					Ingredient{
						Name: "Sugar",
						Amount: "1 lb",
					},
					Ingredient{
						Name: "Butter",
						Amount: "1 stick (1/4 Cup)",
					},
				},
			},
		},
		InstructionParts: []InstructionSection{
			InstructionSection{
				Title: "Frosting",
				Instructions: "MAKE THE FROSTING!",
			},
			InstructionSection{
				Title: "Eating",
				Instructions: "Savor the food.",
			},
		},
	}

	// Check title of the recipe
	if newRecipe.Title != "My new recipe" {
		t.Error("Expected \"My new recipe\", got ", newRecipe.Title)
	}

	// Check the image paths array
	if newRecipe.Images[0] != "image1.jpg" {
		t.Error("Expected the first image to be \"image1.jpg\", got ", newRecipe.Images[0])
	}

	// Check the ingredient list title
	if newRecipe.IngredientLists[0].Title != "Frosting" {
		t.Error("Expected the first ingredient list's name to be \"Frosting\", got ", newRecipe.IngredientLists[0].Title)
	}

	// Check the amount of the second ingredient in the first ingredient list
	if newRecipe.IngredientLists[0].Ingredients[1].Amount != "1 stick (1/4 Cup)" {
		t.Error("Expected the first ingredient list's second ingredient amount to be \"1 stick (1/4 Cup)\", got ", newRecipe.IngredientLists[0].Ingredients[1].Amount)
	}

	// Check the instruction sections
	if newRecipe.InstructionParts[1].Instructions != "Savor the food." {
		t.Error("Expected the instructions for the second instruction part to be \"Savor the food.\", got ", newRecipe.InstructionParts[1].Instructions)
	}
}

// Test ability to Unmarshall the whole struct from JSON
func TestGetRecipeFromJSON(t *testing.T) {
	var newRecipe Recipe
	newRecipe, err := getRecipeFromJSON([]byte(`
		{
			"Title":"My new recipe 2",
			"Source":"Brinker Family 2",
			"Yield":
			"7 servings",
			"Description":"It's really good",
			"Images": ["image1.jpg", "image2.jpg"],
			"IngredientLists": [
				{
					"Title": "Frosting",
					"Ingredients": [
						{
							"Name": "Sugar",
							"Amount": "1 lb"
						},
						{
							"Name": "Butter",
							"Amount": "1 stick (1/4 Cup)"
						}
					]
				}
			],
			"InstructionParts": [
				{
					"Title": "Frosting",
					"Instructions": "MAKE THE FROSTING!"
				},
				{
					"Title": "Eating",
					"Instructions": "Savor the food."
				}
			]
		}`))
	
	// Check for general JSON parse error
	if err != nil {
		t.Error("Error from getRecipeFromJSON: ", err)
	}

	// Check title of the recipe
	if newRecipe.Title != "My new recipe 2" {
		t.Error("Expected \"My new recipe 2\", got ", newRecipe.Title)
	}

	// Check the image paths array
	if newRecipe.Images[0] != "image1.jpg" {
		t.Error("Expected the first image for JSON-Unmarshalled recipe to be \"image1.jpg\", got ", newRecipe.Images[0])
	}

	// Check the ingredient list title
	if newRecipe.IngredientLists[0].Title != "Frosting" {
		t.Error("Expected the first ingredient list's name to be \"Frosting\", got ", newRecipe.IngredientLists[0].Title)
	}

	// Check the amount of the second ingredient in the first ingredient list
	if newRecipe.IngredientLists[0].Ingredients[1].Amount != "1 stick (1/4 Cup)" {
		t.Error("Expected the first ingredient list's second ingredient amount to be \"1 stick (1/4 Cup)\", got ", newRecipe.IngredientLists[0].Ingredients[1].Amount)
	}

	// Check the instruction sections
	if newRecipe.InstructionParts[1].Instructions != "Savor the food." {
		t.Error("Expected the instructions for the second instruction part to be \"Savor the food.\", got ", newRecipe.InstructionParts[1].Instructions)
	}
}

func TestSaveRecipe(t *testing.T) {
	id, err := SaveRecipe([]byte(`
		{
			"Title":"My new recipe 2",
			"Source":"Brinker Family 2",
			"Yield":
			"7 servings",
			"Description":"It's really good",
			"Images": ["image1.jpg", "image2.jpg"],
			"IngredientLists": [
				{
					"Title": "Frosting",
					"Ingredients": [
						{
							"Name": "Sugar",
							"Amount": "1 lb"
						},
						{
							"Name": "Butter",
							"Amount": "1 stick (1/4 Cup)"
						}
					]
				}
			],
			"InstructionParts": [
				{
					"Title": "Frosting",
					"Instructions": "MAKE THE FROSTING!"
				},
				{
					"Title": "Eating",
					"Instructions": "Savor the food."
				}
			]
		}`))
	if err != nil {
		t.Error("Expected nil error, got non-nil error") 
	}

	recipe, err := GetRecipeFromId(id)
	if err != nil {
		t.Error("Expected nil error, got non-nil error")
	}

	if recipe.Title != "My new recipe 2" {
		t.Error("Expected title to be \"My new recipe 2\", got ", recipe.Title)
	}
}