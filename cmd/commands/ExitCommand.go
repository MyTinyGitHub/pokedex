package commands

import (
	"fmt"
	"os"
	"pokedexcli/internal/pokeconfig"
)

func ExitCommand(config *pokeconfig.Config, input []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
