package commands

import (
	"encoding/json"
	"fmt"
	"pokedexcli/internal"
	"pokedexcli/internal/pokeconfig"
)

func Explore(c *pokeconfig.Config, input []string) error {
  area := input[1]
  _, ok := c.Cache.Get(area)
  if !ok { 
    fullUrl := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%v", area)
    data, err := internal.GetFromUrl(fullUrl)
    if err != nil {
      return err
    }

    c.Cache.Add(area, data)
  }

  data, _ := c.Cache.Get(area)

  var explore ExplorePokeApi

  err := json.Unmarshal(data, &explore)
  if err != nil {
    fmt.Println(err)
    return err
  }
  
  encounters := explore.PokemonEncounters
  fmt.Println("Found Pokemon: ")
  for _, encounter := range encounters {
    fmt.Printf(" - %v\n", encounter.Pokemon.Name)
  }

  return nil
}

type ExplorePokeApi struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}
