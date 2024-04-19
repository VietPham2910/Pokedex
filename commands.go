package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func commandHelp() error {
	fmt.Print(
`
Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
`,
	) 
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}