package cmd

import (
	"fmt"
	"math/rand"

	"github.com/VietPham2910/Pokedex/internal/pokeapi"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0{
		return fmt.Errorf("argument missing: Pokemon name")
	}
	pokemon, err := cfg.httpClient.FetchPokemon(pokeapi.PokemonUrl + args[0])
	if err != nil{
		return fmt.Errorf("poke Api fetching error: %v", err)
	}

	res := rand.Intn(pokemon.BaseExperience)
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if _, ok := cfg.pokedex.Pokemons[pokemon.Name]; ok{
		fmt.Printf("%s is already caught!\n", pokemon.Name)
		return nil
	}
	if res > 40 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)
	fmt.Println("You may now inspect it with the inspect command.")
	cfg.pokedex.Add(pokemon)
	return nil
}