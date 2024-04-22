package main

import (
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) == 0{
		return fmt.Errorf("argument missing: Pokemon name")
	}

	name := args[0]
	if _, ok := cfg.pokedex.Pokemons[name]; !ok{
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Println(cfg.pokedex.GetInfo(name))

	return nil
}