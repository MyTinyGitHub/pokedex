package commands

import (
	"fmt"
	"pokedexcli/internal/pokeconfig"
)

func Inspect(c* pokeconfig.Config, input []string) error {
  pokemonName := input[1]
  data, ok := c.CaughtPokemon[pokemonName]

  if !ok {
    fmt.Println("You have not caught that pokemon yet")
    return nil
  }

  data.WriteStats()

  return nil
}
