package main

import "fmt"

func commandHelp(cfg *config) error {
	fmt.Print(
`
Welcome to the Pokedex!
Usage:

`	)
	for key, command := range getCommands() {
		fmt.Printf("%v: %v\n", key, command.description)
	}
	fmt.Print("\n")
	return nil
}