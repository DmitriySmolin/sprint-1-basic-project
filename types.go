package main

// Stock — остатки ингредиентов
type Stock map[string]int

// Recipe — рецепт напитка
type Recipe struct {
	Price       int
	Ingredients map[string]int
	Steps       []string
}

// Stats — статистика продаж
type Stats struct {
	Orders  int
	Revenue int
}

// State — сохраняемое состояние
type State struct {
	Stock Stock `json:"stock"`
	Stats Stats `json:"stats"`
}
