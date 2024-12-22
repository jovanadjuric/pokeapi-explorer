package pokeapi

import (
	"net/http"
	"time"

	"github.com/jovanadjuric/pokedex/internal/pokecache"
)

// Client -
type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

// NewClient -
func NewClient(timeout time.Duration, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(5 * cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
