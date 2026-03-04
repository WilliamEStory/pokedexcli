package main

import "fmt"

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		fmt.Println("Please specify a location name")
		return nil
	}

	locName := args[0]

	pokemonResp, err := cfg.pokeapiClient.ListPokemonForLocation(locName)
	if err != nil {
		return err
	}

	for _, pokemon := range pokemonResp.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}

	return nil
}
