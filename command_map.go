package main

import (
	"fmt"

	"github.com/thetramp22/pokedexcli/internal/pokedata"
)

func commandMap(cfg *config, args ...string) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.nextLocationsURL != nil {
		url = *cfg.nextLocationsURL
	}
	if cfg.previousLocationsURL != nil && cfg.nextLocationsURL == nil {
		fmt.Println("You're on the last page")
		return nil
	}

	locationAreas, err := pokedata.GetLocationAreas(url, cfg.cache)
	if err != nil {
		return err
	}
	for _, result := range locationAreas.Results {
		fmt.Println(result.Name)
	}
	cfg.nextLocationsURL = locationAreas.Next
	cfg.previousLocationsURL = locationAreas.Previous
	return nil
}

func commandMapb(cfg *config, args ...string) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.previousLocationsURL != nil {
		url = *cfg.previousLocationsURL
	}
	if cfg.previousLocationsURL == nil {
		fmt.Println("You're on the first page")
		return nil
	}

	locationAreas, err := pokedata.GetLocationAreas(url, cfg.cache)
	if err != nil {
		return err
	}
	for _, result := range locationAreas.Results {
		fmt.Println(result.Name)
	}
	cfg.nextLocationsURL = locationAreas.Next
	cfg.previousLocationsURL = locationAreas.Previous
	return nil
}
