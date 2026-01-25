package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const stateFile = "coffee_state.json"

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

func (m *Machine) Load() error {
	data, err := os.ReadFile(stateFile)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}
	var state State
	if err := json.Unmarshal(data, &state); err != nil {
		return err
	}
	m.Stock = state.Stock
	m.Stats = state.Stats
	return nil
}

func (m *Machine) Save() error {
	state := State{Stock: m.Stock, Stats: m.Stats}
	data, err := json.Marshal(state)
	if err != nil {
		return err
	}
	return os.WriteFile(stateFile, data, 0644)
}
func (m *Machine) PrintStats() {
	fmt.Printf("orders=%d revenue=%d\n", m.Stats.Orders, m.Stats.Revenue)
}
