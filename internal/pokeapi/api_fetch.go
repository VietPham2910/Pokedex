package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func FetchData[T any](client *Client, url string) (T ,error){
	body, ok := client.cache.Get(url)
	var v T 
	if !ok {
		res, err := client.httpClient.Get(url)
		if err != nil {
			return v, fmt.Errorf("url fetching error: %v", err)
		}
	
		body, err = io.ReadAll(res.Body)
		if err != nil {
			return v, fmt.Errorf("couldn't read response body: %v", err)
		}

		client.cache.Add(url, body)
	}

	err := json.Unmarshal(body, &v)
	if err != nil{
		if e, ok := err.(*json.SyntaxError); ok {
			return v, fmt.Errorf("syntax error at byte offset %d", e.Offset)
		}
		return v, fmt.Errorf("couldn't convert json to %T object: %v",v, err)
	}

	return v, nil
}

func (client *Client)FetchLocAreas(url string) (LocationAreas ,error){
	return FetchData[LocationAreas](client, url)
}

func (client *Client)FetchLocDetail(url string) (LocationDetail ,error){
	return FetchData[LocationDetail](client, url)
}

func (client *Client)FetchPokemon(url string) (Pokemon ,error){
	return FetchData[Pokemon](client, url)

}
