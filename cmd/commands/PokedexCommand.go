package commands

import (
	"fmt"
	"pokedexcli/internal/pokeconfig"
)

func Pokedex(c *pokeconfig.Config, input []string) error {
  fmt.Println("Your Pokedex: ")

  for key := range c.CaughtPokemon {
    fmt.Printf("  - %v \n", key)
  }

  return nil
}
