package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FetchPoke(url string) (*LocationAreas ,error){
	fmt.Printf("Fetching %v...\n", url)
	
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("url fetching error: %v", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("couldn't read response body: %v", err)
	}

	locAreas := LocationAreas{} 
	err = json.Unmarshal(body, &locAreas)
	if err != nil{
		return nil, fmt.Errorf("couldn't convert json to LocationAreas object: %v", err)
	}

	return &locAreas, nil
}