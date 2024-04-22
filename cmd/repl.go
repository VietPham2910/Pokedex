package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/VietPham2910/Pokedex/internal"
	"github.com/VietPham2910/Pokedex/internal/pokeapi"
)

type config struct {
	httpClient pokeapi.Client
	nextLocationUrl     string
	previousLocationUrl string
	pokedex internal.Pokedex
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func NewConfig() *config{
	return &config{
		httpClient: *pokeapi.NewClient(time.Second * 5, time.Minute * 5),
		nextLocationUrl: pokeapi.LocationUrl,
		pokedex: internal.Pokedex{
			Pokemons: make(map[string]pokeapi.Pokemon),
		},
	}
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

func StartRepl(){
	cfg := NewConfig()
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

