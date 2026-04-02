package main

import (
	"fmt"

	"github.com/thetramp22/pokedexcli/internal/pokedata"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("Please enter a valid area")
	}
	area := args[1]
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%v/", area)
	locationArea, err := pokedata.GetLocationArea(url, cfg.cache)
	if err != nil {
		return fmt.Errorf("%v is not a valid area", area)
	}
	fmt.Printf("Exploring %v\n", locationArea.Name)
	fmt.Println("Found Pokemon:")
	for _, result := range locationArea.PokemonEncounters {
		fmt.Printf("- %v\n", result.Pokemon.Name)
	}
	return nil
}
