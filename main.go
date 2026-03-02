package main

import (
	"strings"
)

func main() {
	startRepl()
}

func cleanInput(text string) []string {
	cleaned := strings.TrimSpace(text)

	return strings.Fields(cleaned)
}
