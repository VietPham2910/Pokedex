package main

import (
	"time"

	"github.com/VietPham2910/Pokedex/internal/pokeapi"
)

type config struct {
	httpClient pokeapi.Client
	nextLocationUrl     string
	previousLocationUrl string
	pokedex Pokedex
}

func main() {
	cfg := &config{
		httpClient: *pokeapi.NewClient(time.Second * 5, time.Minute * 5),
		nextLocationUrl: pokeapi.LocationUrl,
		pokedex: Pokedex{
			Pokemons: make(map[string]pokeapi.Pokemon),
		},
	}
	startRepl(cfg)
}