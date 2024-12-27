package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// GetPokemon -
func (c *Client) GetPokemon(pageURL *string, pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locationResp := Pokemon{}

		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return Pokemon{}, err
		}

		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	locationResp := Pokemon{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, dat)

	return locationResp, nil
}
