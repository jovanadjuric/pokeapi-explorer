package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	pokeapi "github.com/jovanadjuric/pokedex/internal/pokeapi"
)

type cliCommand struct {
	callback    func(*config) error
	name        string
	description string
}

type config struct {
	nextLocationsURL *string
	prevLocationsURL *string
	pokeapiClient    pokeapi.Client
}

func startRepl(cfg *config) {
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
					command.callback(cfg)
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
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
	}
}
