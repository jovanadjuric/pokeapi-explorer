package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, params ...string) error {
	if len(params) == 0 {
		return errors.New("missing pokemon name")
	}

	pokemonName := params[0]

	fmt.Println("Throwing a Pokeball at " + pokemonName + "...")

	pokemon, err := cfg.pokeapiClient.GetPokemon(cfg.nextLocationsURL, pokemonName)
	if err != nil {
		return err
	}

	randomNumber := rand.Intn(300)
	isCaught := randomNumber < pokemon.BaseExperience

	if isCaught {
		fmt.Println(pokemonName + " was caught!")
		cfg.pokemons[pokemonName] = pokemon
		return nil
	}

	fmt.Println(pokemonName + " escaped!")
	return nil
}
