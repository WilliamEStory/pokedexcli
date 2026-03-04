package main

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.pokedex) == 0 {
		println("Your pokedex is empty. Try catching some pokemon!")
		return nil
	}

	println("Your Pokedex:")
	for name := range cfg.pokedex {
		println(" - ", name)
	}

	return nil
}
