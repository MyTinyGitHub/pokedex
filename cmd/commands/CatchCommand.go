package commands

import (
	"encoding/json"
	"fmt"
	"pokedexcli/internal"
	"pokedexcli/internal/pokeconfig"
)


func Catch(c *pokeconfig.Config, input []string) error {
  pokemonName := input[1]


  fmt.Printf("Throwing a Pokeball at %v...\n", pokemonName)

  fullUrl := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%v", pokemonName)
  data, err := internal.GetFromUrl(fullUrl)
  if err != nil {
    fmt.Println(err)
    return err
  }

  var pokemon internal.Pokemon
  err = json.Unmarshal(data, &pokemon)
  if err != nil {
    fmt.Println(err)
    return err
  }

  caught := pokemon.TryCatch()
  if caught {
    c.CaughtPokemon[pokemonName] = pokemon
  }

  return nil
}


