package repl

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Location struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap() error {
	url := "https://pokeapi.co/api/v2/location"

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error fetching data: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-200 response: %d", resp.StatusCode)
	}

	dec := json.NewDecoder(resp.Body)
	var locations Location

	if err := dec.Decode(&locations); err != nil {
		return fmt.Errorf("error decoding JSON: %v", err)
	}

	dec.Decode(&locations)

	for _, location := range locations.Results {
		fmt.Printf("%s\n", location.Name+"-area")
	}

	return nil
}
