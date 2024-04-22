package internal

import (
	"fmt"
	"strings"

	"github.com/VietPham2910/Pokedex/internal/pokeapi"
)

type Pokedex struct{
	Pokemons map[string]pokeapi.Pokemon
}

func (pokedex *Pokedex) Add (pokemon pokeapi.Pokemon){
	if _, ok := pokedex.Pokemons[pokemon.Name]; !ok{
		pokedex.Pokemons[pokemon.Name] = pokemon
	}
}

func (pokedex *Pokedex) GetStats (name string) string{
	var sb strings.Builder
	for _, stat := range pokedex.Pokemons[name].Stats {
		sb.WriteString(fmt.Sprintf("\n -%s: %d", stat.Stat.Name, stat.BaseStat))
	}

	return sb.String()
}

func (pokedex *Pokedex) GetTypes (name string) string{
	var sb strings.Builder
	for _, pokemonType := range pokedex.Pokemons[name].Types {
		sb.WriteString(fmt.Sprintf("\n -%s", pokemonType.Type.Name))
	}
	
	return sb.String()
}

func (pokedex *Pokedex) GetInfo (name string) string{
	return fmt.Sprintf(
`Name: %s
Height: %d
Weight: %d
Stats:%v
Types:%v`,  name,
			pokedex.Pokemons[name].Height,
			pokedex.Pokemons[name].Weight, 
			pokedex.GetStats(name), 
			pokedex.GetTypes(name),
)}
