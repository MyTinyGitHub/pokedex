package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pokedexcli/internal/pokeconfig"
)

type LocationArea struct {
	Count    int               `json:"count"`
	Next     string            `json:"next"`
	Previous string            `json:"previous"`
	Location []LocationResults `json:"results"`
}

type LocationResults struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func Map(config *pokeconfig.Config, input []string) error {
	locations, err := getLocationAreas(config.Next, config)
	if err != nil {
		return err
	}

	config.Next = locations.Next
	config.Previous = locations.Previous

	for _, location := range locations.Location {
		fmt.Println(location.Name)
	}

	return nil
}

func MapBack(config *pokeconfig.Config, input []string) error {
	if config.Previous == "" {
		fmt.Println("You're on the first page")
		return nil
	}

	locations, err := getLocationAreas(config.Previous, config)
	if err != nil {
		return err
	}

	config.Next = locations.Next
	config.Previous = locations.Previous

	for _, location := range locations.Location {
		fmt.Println(location.Name)
	}

	return nil
}

func getDataFromUrl(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	return io.ReadAll(res.Body)
}

func getData(url string, c *pokeconfig.Config) ([]byte, error) {
	value, ok := c.Cache.Get(url)
	if !ok {
		data, err := getDataFromUrl(url)
		if err != nil {
			return nil, err
		}

		c.Cache.Add(url, data)

		return data, nil
	}
	return value, nil
}

func getLocationAreas(url string, c *pokeconfig.Config) (LocationArea, error) {
	value, err := getData(url, c)
	if err != nil {
		return LocationArea{}, err
	}

	var locations LocationArea

	err = json.Unmarshal(value, &locations)
	if err != nil {
		return LocationArea{}, fmt.Errorf("unable to unmarshal: %v", err)
	}

	return locations, nil
}
