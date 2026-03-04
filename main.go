package main

import (
	"time"

	"github.com/williamestory/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokedex:       make(map[string]pokeapi.RespPokemon),
	}

	startRepl(cfg)
}
