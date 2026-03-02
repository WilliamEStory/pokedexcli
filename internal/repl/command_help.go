package repl

import (
	"fmt"
)

func commandHelp() error {
	commands := initCommands()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")

	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	return nil
}
