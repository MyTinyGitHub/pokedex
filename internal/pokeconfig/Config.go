package pokeconfig

import (
	"pokedexcli/internal/pokecache"
)

type Config struct {
	Next     string
	Previous string
	Cache    pokecache.Cache
	Registry map[string]CliCommand
}
