package main

import "fmt"

func main() {
	fmt.Println(NewMachine().Recipes["espresso"].Price)
}
