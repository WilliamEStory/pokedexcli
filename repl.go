package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/williamestory/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *config) {
	scan := bufio.NewScanner(os.Stdin)

	for {
		commands := initCommands()
		fmt.Print("Pokedex > ")
		if !scan.Scan() {
			fmt.Println("Error reading input:", scan.Err())
			continue
		}

		input := scan.Text()
		cleanedInput := cleanInput(input)

		if len(cleanedInput) == 0 {
			fmt.Println("Specify a command")
		} else {
			cmdName := cleanedInput[0]
			cmd, exists := commands[cmdName]
			if !exists {
				fmt.Println("Unknown command")
			} else {
				err := cmd.callback(cfg)
				if err != nil {
					fmt.Println("Error executing command:", err)
				}
			}
		}
	}
}

func initCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays a list of pokemon locations.",
			callback:    commandMapF,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous page of pokemon locations.",
			callback:    commandMapB,
		},
	}
}

func cleanInput(text string) []string {
	cleaned := strings.TrimSpace(text)

	return strings.Fields(cleaned)
}
