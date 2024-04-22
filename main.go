package main

import (
	"github.com/VietPham2910/Pokedex/internal/pokeapi"
)

func main() {
	cfg := &config{
		pokeEndpointUrl: pokeapi.EndpointUrl,
		nextLocationUrl: pokeapi.EndpointUrl,
		previousLocationUrl: "",
	}
	startRepl(cfg)
}