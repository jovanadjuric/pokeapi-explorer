package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, params ...string) error {
	if len(params) == 0 {
		return errors.New("missing location name")
	}

	locationName := params[0]

	fmt.Println("Exploring " + locationName + "...")

	locationsResp, err := cfg.pokeapiClient.ListPokemons(cfg.nextLocationsURL, locationName)
	if err != nil {
		return err
	}

	for _, enc := range locationsResp.PokemonEncounters {
		fmt.Println("- " + enc.Pokemon.Name)
	}
	return nil
}
