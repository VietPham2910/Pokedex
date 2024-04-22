package main

import (
	"fmt"

	"github.com/VietPham2910/Pokedex/internal/pokeapi"
)

type config struct {
	pokeEndpointUrl     string
	nextLocationUrl     string
	previousLocationUrl string
}

func commandMapf(cfg *config) error {
	locations, err := pokeapi.FetchPoke(cfg.nextLocationUrl)
	if err != nil{
		return fmt.Errorf("poke Api fetching error: %v", err)
	}
	
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	cfg.nextLocationUrl = locations.Next
	if locations.Previous == nil {
		cfg.previousLocationUrl = ""
	} else{
		cfg.previousLocationUrl = *locations.Previous
	}

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.previousLocationUrl == ""{
		return fmt.Errorf("empty previous path")
	}
	locations, err := pokeapi.FetchPoke(cfg.previousLocationUrl)
	if err != nil{
		return fmt.Errorf("poke Api fetching error: %v", err)
	}
	
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	cfg.nextLocationUrl = locations.Next
	if locations.Previous == nil {
		cfg.previousLocationUrl = ""
	} else{
		cfg.previousLocationUrl = *locations.Previous
	}

	return nil
}