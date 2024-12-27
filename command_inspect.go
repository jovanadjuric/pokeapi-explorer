package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, params ...string) error {
	if len(params) == 0 {
		return errors.New("missing location name")
	}

	pokemonName := params[0]

	pokemon, ok := cfg.pokemons[pokemonName]

	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %v\n", *pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", *stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, typ := range pokemon.Types {
		fmt.Printf("  - %s\n", *typ.Type.Name)
	}

	return nil
}
