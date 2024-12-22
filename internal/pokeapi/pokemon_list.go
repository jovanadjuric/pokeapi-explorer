package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListPokemons -
func (c *Client) ListPokemons(pageURL *string, locationName string) (RespShallowLocation, error) {
	url := baseURL + "/location-area/" + locationName
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locationResp := RespShallowLocation{}

		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return RespShallowLocation{}, err
		}

		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocation{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocation{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocation{}, err
	}

	locationResp := RespShallowLocation{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return RespShallowLocation{}, err
	}

	c.cache.Add(url, dat)

	return locationResp, nil
}
