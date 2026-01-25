package main

var recipes = map[string]Recipe{
	"espresso": {
		Price:       150,
		Ingredients: map[string]int{"beans": 8, "water": 30},
		Steps:       []string{"grind", "tamp", "brew 25s"},
	},
	"americano": {
		Price:       180,
		Ingredients: map[string]int{"beans": 8, "water": 120},
		Steps:       []string{"grind", "brew 25s", "add water"},
	},
	"latte": {
		Price:       220,
		Ingredients: map[string]int{"beans": 8, "water": 30, "milk": 150},
		Steps:       []string{"grind", "brew 25s", "steam milk", "mix"},
	},
}

var recipeOrder = []string{"espresso", "americano", "latte"}

var validIngredients = map[string]bool{
	"beans": true, "water": true, "milk": true, "sugar": true,
}

var ingredientOrder = []string{"beans", "water", "milk", "sugar"}