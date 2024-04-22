package cmd

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.pokedex.Pokemons {
		fmt.Print(" - ")
		fmt.Println(pokemon.Name)
	}

	return nil
}