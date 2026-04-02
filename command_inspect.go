package main

import "fmt"

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("Please enter a valid Pokemon")
	}
	pokemon := args[1]
	if _, ok := cfg.userDex[pokemon]; !ok {
		return fmt.Errorf("You have not caught that Pokemon")
	}
	currentPokemon := cfg.userDex[pokemon]
	fmt.Printf("Name: %v\n", currentPokemon.Name)
	fmt.Printf("Height: %v\n", currentPokemon.Height)
	fmt.Printf("Weight: %v\n", currentPokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range currentPokemon.Stats {
		fmt.Printf("  -%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pType := range currentPokemon.Types {
		fmt.Printf("  -%v\n", pType.Type.Name)
	}
	return nil
}
