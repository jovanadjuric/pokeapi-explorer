package main

import (
	"fmt"
)

func commandPokedex(cfg *config, params ...string) error {
	pokemons := cfg.pokemons

	fmt.Println("Your Pokedex:")
	for _, p := range pokemons {
		fmt.Println(" - " + *p.Name)
	}

	return nil
}
