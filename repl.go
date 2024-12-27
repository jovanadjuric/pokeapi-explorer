package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	pokeapi "github.com/jovanadjuric/pokedex/internal/pokeapi"
)

type cliCommand struct {
	callback    func(*config, ...string) error
	name        string
	description string
}

type config struct {
	nextLocationsURL *string
	prevLocationsURL *string
	pokemons         map[string]pokeapi.Pokemon
	pokeapiClient    pokeapi.Client
}

func startRepl(cfg *config) {
	commands := getCommands()

	scanner := bufio.NewScanner(os.Stdin)

	executed := false
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() && executed {
			break
		}

		executed = true

		text := scanner.Text()
		sanitized := strings.Fields(strings.ToLower(text))

		if len(sanitized) == 0 {
			executed = false
			continue
		}

		found := false
		for k, command := range commands {
			if k == sanitized[0] {
				found = true
				if len(sanitized) > 1 {
					command.callback(cfg, sanitized[1:]...)
					break
				}

				err := command.callback(cfg)
				if err != nil {
					fmt.Println(err)
				}
				break
			}
		}

		if !found {
			fmt.Println("Unknown command")
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
		"explore": {
			name:        "explore <location_name>",
			description: "Explore a location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <location_name>",
			description: "Catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "Inspect a pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Display a list of all caught pokemons",
			callback:    commandPokedex,
		},
	}
}
