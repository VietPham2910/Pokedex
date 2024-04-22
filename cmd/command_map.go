package cmd

import (
	"fmt"
)

func commandMapf(cfg *config, args ...string) error {
	locations, err := cfg.httpClient.FetchLocAreas(cfg.nextLocationUrl)
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

func commandMapb(cfg *config, args ...string) error {
	if cfg.previousLocationUrl == ""{
		return fmt.Errorf("empty previous path")
	}
	locations, err := cfg.httpClient.FetchLocAreas(cfg.previousLocationUrl)
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