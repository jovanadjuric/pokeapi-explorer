package main

import (
	"time"

	pokeApi "github.com/jovanadjuric/pokedex/internal/pokeApi"
)

func main() {
	pokeClient := pokeApi.NewClient(5 * time.Second)

	cfg := &config{
		pokeApiClient: pokeClient,
	}

	startRepl(cfg)
}
