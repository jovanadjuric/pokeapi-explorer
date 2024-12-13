package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type pokeApiResponse struct {
	Next     string
	Previous string
	Results  []pokeApiLocation
	Count    int
}

type pokeApiLocation struct {
	Name string
	Url  string
}

var page int = 0

func commandMap() error {
	limit := 20
	offset := 20 * page

	res, err := http.Get("https://pokeapi.co/api/v2/location-area" + "?limit=" + strconv.Itoa(limit) + "&offset=" + strconv.Itoa(offset))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	page++

	decoder := json.NewDecoder(res.Body)
	var formattedResponse pokeApiResponse

	err = decoder.Decode(&formattedResponse)
	if err != nil {
		return err
	}

	for _, v := range formattedResponse.Results {
		fmt.Println(v.Name)
	}

	return nil
}
