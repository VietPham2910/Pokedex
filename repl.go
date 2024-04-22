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
	callback    func(*config) error
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
				err := command.callback(cfg)
				if err != nil{
					fmt.Println(err)
				}
			} else{
				fmt.Print("Command not found!")
			}
			break
		}
	}
}

