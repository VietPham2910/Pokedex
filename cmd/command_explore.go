package cmd

import (
	"fmt"

	"github.com/VietPham2910/Pokedex/internal/pokeapi"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0{
		return fmt.Errorf("argument missing: location name")
	}
	fmt.Printf("Exploring %s...\n", args[0])
	locationDetail, err := cfg.httpClient.FetchLocDetail(pokeapi.LocationUrl + args[0])
	if err != nil{
		return fmt.Errorf("poke Api fetching error: %v", err)
	}
	
	fmt.Println("Found Pokemon:")
	for _, pokemon := range locationDetail.PokemonEncounters {
		fmt.Print(" - ")
		fmt.Println(pokemon.Pokemon.Name)
	}

	return nil
}