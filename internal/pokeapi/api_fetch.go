package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (client *Client)FetchPoke(url string) (*LocationAreas ,error){
	fmt.Printf("Fetching %v...\n", url)

	body, ok := client.cache.Get(url)
	if !ok {
		res, err := client.httpClient.Get(url)
		if err != nil {
			return nil, fmt.Errorf("url fetching error: %v", err)
		}
	
		body, err = io.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("couldn't read response body: %v", err)
		}

		client.cache.Add(url, body)
	}

	locAreas := LocationAreas{} 
	err := json.Unmarshal(body, &locAreas)
	if err != nil{
		return nil, fmt.Errorf("couldn't convert json to LocationAreas object: %v", err)
	}

	return &locAreas, nil
}