package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Get a list of all the Pokémon in a given area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Try catching a Pokémon with a given name",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Find and inspect Pokémon info with a given name",
			callback:    commandInspect,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func startRepl(cfg *config){
	scanner := bufio.NewScanner(os.Stdin)
	for{
		fmt.Print("pokedex > ")
		for scanner.Scan(){
			words := cleanInput(scanner.Text())
			if len(words) == 0{
				continue
			}
			if command, ok := getCommands()[words[0]]; ok{
				err := command.callback(cfg, words[1:]...)
				if err != nil{
					fmt.Println(err)
				}	
			} else{
				fmt.Println("Command not found!")
			}
			break
		}
	}
}

