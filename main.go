package main

import (
	"time"

	pokeapi "github.com/jovanadjuric/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)

	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
