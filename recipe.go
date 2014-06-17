package main

import (
	"encoding/json"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type Recipe struct {
	_id string
	Title string
	Source string
	Yield string
	Description string
	Images []string
	IngredientLists []IngredientSection
	InstructionParts []InstructionSection
}

type IngredientSection struct {
	Title string
	Ingredients []Ingredient
}

type Ingredient struct {
	Name string
	Amount string
}

type InstructionSection struct {
	Title string
	Instructions string
}

func getRecipeFromJSON( b []byte ) (Recipe, error) {
	var r Recipe
	err := json.Unmarshal(b, &r)
	if err != nil {
		return Recipe{}, err
	}
	return r, nil
}

func SaveRecipe( b []byte ) (bson.ObjectId, error) {
	var r Recipe
	r, err := getRecipeFromJSON(b)
	if err != nil {
		return "", err
	}

	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		return "", err
	}
	defer session.Close()

	objID := bson.NewObjectId()

	c := session.DB("tag730").C("recipes")
	_, err = c.UpsertId(objID,&r)
	if err != nil {
		return "", err
	}
	return objID, nil
}

func GetRecipeFromId( id bson.ObjectId ) (Recipe, error) {
	var r Recipe

	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		return Recipe{}, err
	}
	defer session.Close()

	c := session.DB("tag730").C("recipes")
	err = c.FindId(id).One(&r)
	if err != nil {
		return Recipe{}, err
	}
	return r, nil
}

