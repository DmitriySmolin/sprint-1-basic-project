package main

import "github.com/k0kubun/pp/v3"

func main() {
	machine := NewMachine()

	// fmt.Println(machine.Recipes["espresso"].Price)

	// machine.GetStock()

	// machine.AddStock("beans", 2)
	// pp.Println(machine)

	machine.SetStock("water", 3)
	// pp.Println(machine)

	machine.AddStock("beans", 10)
	machine.AddStock("water", 50)

	// Варим espresso с оплатой 150
	err := machine.Brew("espresso", 150)
	if err != nil {
		pp.Println("Ошибка:", err)
	}

}
