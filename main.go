package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands map[string]cliCommand

func main() {
	initCommands()
	scan := bufio.NewScanner(os.Stdin)

	for {
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

func cleanInput(text string) []string {
	cleaned := strings.TrimSpace(text)

	return strings.Fields(cleaned)
}

func initCommands() {
	commands = map[string]cliCommand{
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
