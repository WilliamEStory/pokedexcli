package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		fmt.Println("Please specify a pokemon name")
		return nil
	}

	pokemonName := args[0]

	pokemonResp, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	caught := float32(pokemonResp.BaseExperience)/100.0 <= rand.Float32()+0.5

	if caught {
		fmt.Println("You caught", pokemonName)
		_, exists := cfg.pokedex[pokemonName]
		if !exists {
			cfg.pokedex[pokemonName] = pokemonResp
		} else {
			fmt.Println("You already have", pokemonName)
		}
	} else {
		fmt.Println(pokemonName, "escaped")
	}

	return nil
}
