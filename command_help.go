package main

import "fmt"

func commandHelp(cfg *config, _ ...string) error {
	commands := getCommands()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")

	for _, command := range commands {
		fmt.Println(command.name + ": " + command.description)
	}

	return nil
}
