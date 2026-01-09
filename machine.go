package main

type Machine struct {
	Stock   Stock
	Recipes map[string]Recipe
	Stats   Stats
}

func NewMachine() *Machine {
	m := &Machine{
		Stock:   Stock{"beans": 0, "water": 0, "milk": 0, "sugar": 0},
		Recipes: recipes,
	}
	return m
}
