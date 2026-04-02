package main

import (
	"fmt"

	"github.com/thetramp22/pokedexcli/internal/pokedata"
)

func commandMap(cfg *config, args ...string) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.Next != nil {
		url = *cfg.Next
	}
	if cfg.Previous != nil && cfg.Next == nil {
		fmt.Println("You're on the last page")
		return nil
	}

	locationAreas, err := pokedata.GetLocationAreas(url, cfg.Cache)
	if err != nil {
		return err
	}
	for _, result := range locationAreas.Results {
		fmt.Println(result.Name)
	}
	cfg.Next = locationAreas.Next
	cfg.Previous = locationAreas.Previous
	return nil
}

func commandMapb(cfg *config, args ...string) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.Previous != nil {
		url = *cfg.Previous
	}
	if cfg.Previous == nil {
		fmt.Println("You're on the first page")
		return nil
	}

	locationAreas, err := pokedata.GetLocationAreas(url, cfg.Cache)
	if err != nil {
		return err
	}
	for _, result := range locationAreas.Results {
		fmt.Println(result.Name)
	}
	cfg.Next = locationAreas.Next
	cfg.Previous = locationAreas.Previous
	return nil
}
