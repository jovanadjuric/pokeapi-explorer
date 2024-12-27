package main

import (
	"time"

	pokeapi "github.com/jovanadjuric/pokeapi-explorer/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)

	cfg := &config{
		pokeapiClient: pokeClient,
		pokemons:      make(map[string]pokeapi.Pokemon),
	}

	startRepl(cfg)
}
