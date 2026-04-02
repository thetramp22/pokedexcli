package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.UserDex) == 0 {
		fmt.Println("Your Pokedex is empty")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.UserDex {
		fmt.Printf("  -%v\n", pokemon.Name)
	}
	return nil
}
