package commands

import (
	"fmt"
	"pokedexcli/internal/pokeconfig"
)

func Help(config *pokeconfig.Config, input []string) error {
	for _, value := range config.Registry {
		fmt.Printf("%v: %v\n", value.Name, value.Description)
	}

	return nil
}
