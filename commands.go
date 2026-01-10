package main

import "fmt"

func (m *Machine) GetStock() {
	for _, ing := range ingredientOrder {
		fmt.Printf("%s=%d\n", ing, m.Stock[ing])
	}
}

func (m *Machine) AddStock(ingredient string, qty int) error {
	if !validIngredients[ingredient] || qty < 1 {
		return fmt.Errorf("некорректные параметры")
	}
	m.Stock[ingredient] += qty
	fmt.Println("ok")
	return nil
}

func (m *Machine) SetStock(ingredient string, qty int) error {
	if !validIngredients[ingredient] || qty < 0 {
		return fmt.Errorf("некорректные параметры")
	}
	m.Stock[ingredient] = qty
	fmt.Println("ok")
	return nil
}
