package main

import (
	"time"

	"github.com/thetramp22/pokedexcli/internal/pokecache"
)

func main() {
	interval := 5 * time.Second
	cfg := config{
		Next:     nil,
		Previous: nil,
		Cache:    pokecache.NewCache(interval),
	}
	startRepl(&cfg)
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

type config struct {
	Next     *string
	Previous *string
	Cache    *pokecache.Cache
}
