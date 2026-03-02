package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func startRepl() {
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
				err := cmd.callback()
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
	}
}
