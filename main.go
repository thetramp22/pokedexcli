package main

import (
	"time"

	"github.com/thetramp22/pokedexcli/internal/pokecache"
	"github.com/thetramp22/pokedexcli/internal/pokedata"
)

func main() {
	interval := 5 * time.Second
	cfg := config{
		cache:   pokecache.NewCache(interval),
		userDex: map[string]pokedata.Pokemon{},
	}
	startRepl(&cfg)
}
