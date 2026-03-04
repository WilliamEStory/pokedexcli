package main

func commandInspect(cfg *config, args ...string) error {
	if len(args) == 0 {
		println("Please specify a pokemon name")
		return nil
	}

	pokemonName := args[0]
	pokemonResp, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	println("Name: ", pokemonResp.Name)
	println("Height: ", pokemonResp.Height)
	println("Stats:")
	for _, stat := range pokemonResp.Stats {
		println("  ", stat.Stat.Name, ":", stat.BaseStat)
	}
	println("Types:")
	for _, t := range pokemonResp.Types {
		println("  ", t.Type.Name)
	}
	return nil
}
