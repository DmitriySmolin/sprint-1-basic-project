package main

import (
	"fmt"
	"os"
	"strconv"
)

func printUsage() {
	fmt.Println("usage")
}
func main() {
	machine := NewMachine()
	if err := machine.Load(); err != nil {
		fmt.Println("ошибка")
		os.Exit(3)
	}
	args := os.Args[1:]

	if len(args) == 0 {
		printUsage()
		os.Exit(2)
	}

	var needSave bool

	switch args[0] {
	case "help":
		printUsage()
		os.Exit(0)

	case "menu":
		machine.Menu()

	case "stock":
		if len(args) < 2 {
			printUsage()
			os.Exit(2)
		}
		switch args[1] {
		case "get":
			machine.GetStock()
		case "add":
			if len(args) < 4 {
				printUsage()
				os.Exit(2)
			}
			qty, err := strconv.Atoi(args[3])
			if err != nil {
				fmt.Println("некорректные параметры")
				os.Exit(1)
			}
			if err := machine.AddStock(args[2], qty); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			needSave = true
		case "set":
			if len(args) < 4 {
				printUsage()
				os.Exit(2)
			}
			qty, err := strconv.Atoi(args[3])
			if err != nil {
				fmt.Println("некорректные параметры")
				os.Exit(1)
			}
			if err := machine.SetStock(args[2], qty); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			needSave = true
		default:
			printUsage()
			os.Exit(2)
		}

	case "brew":
		if len(args) < 4 || args[2] != "--pay" {
			printUsage()
			os.Exit(2)
		}
		payment, err := strconv.Atoi(args[3])
		if err != nil {
			printUsage()
			os.Exit(2)
		}
		if err := machine.Brew(args[1], payment); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		needSave = true

	case "stats":
		machine.PrintStats()

	default:
		printUsage()
		os.Exit(2)
	}

	if needSave {
		if err := machine.Save(); err != nil {
			fmt.Println("ошибка")
			os.Exit(3)
		}
	}
}
