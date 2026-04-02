package main

import (
	"fmt"
	"math/rand"

	"github.com/thetramp22/pokedexcli/internal/pokedata"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("Please enter a valid Pokemon")
	}
	pokemon := args[1]
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%v/", pokemon)
	currentPokemon, err := pokedata.GetPokemon(url, cfg.cache)
	if err != nil {
		return fmt.Errorf("%v is not a valid Pokemon", pokemon)
	}
	if _, ok := cfg.userDex[pokemon]; ok {
		fmt.Printf("You already have a %v in your Pokedex\n", pokemon)
		return nil
	}
	fmt.Printf("Throwing a Pokeball at %v...\n", currentPokemon.Name)
	attempt := rand.Intn(400)
	if attempt < currentPokemon.BaseExperience {
		fmt.Printf("%v escaped!\n", currentPokemon.Name)
		return nil
	}
	fmt.Printf("%v was caught!\n", currentPokemon.Name)
	cfg.userDex[currentPokemon.Name] = currentPokemon
	return nil
}
