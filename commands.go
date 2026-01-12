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

func (m *Machine) Brew(drink string, payment int) error {
	recipe, ok := m.Recipes[drink]

	if !ok {
		return fmt.Errorf("напиток не найден")
	}

	if payment < recipe.Price {
		return fmt.Errorf("недостаточно оплаты")
	}

	for ing, need := range recipe.Ingredients {
		if m.Stock[ing] < need {
			return fmt.Errorf("не хватает ингредиентов")
		}
	}

	for ing, need := range recipe.Ingredients {
		m.Stock[ing] -= need
	}

	m.Stats.Orders++
	m.Stats.Revenue += recipe.Price
	for _, step := range recipe.Steps {
		fmt.Println(step)
	}

	return nil
}
