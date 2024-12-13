package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	callback    func() error
	name        string
	description string
}

func startRepl() {
	commands := getCommands()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			text := scanner.Text()
			sanitized := strings.Fields(strings.ToLower(text))

			found := false
			for _, command := range commands {
				if command.name == sanitized[0] {
					found = true
					command.callback()
					break
				}
			}

			if !found {
				fmt.Println("Unknown command")
			}
		}
	}
}

func cleanInput(text string) []string {
	var words []string

	for _, word := range strings.Split(text, " ") {
		fmt.Println(word)
		trimmedText := strings.TrimSpace(word)
		if trimmedText != "" {
			words = append(words, trimmedText)
		}
	}

	return words
}

func getCommands() map[string]cliCommand {
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
