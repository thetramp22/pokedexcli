package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/thetramp22/pokedexcli/internal/pokedata"
)

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	words := strings.Fields(lowerText)
	return words
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		line := scanner.Text()
		words := cleanInput(line)
		command := words[0]
		param1 := ""
		if len(words) == 2 {
			param1 = words[1]
		}
		commands := getCommands()
		if _, ok := commands[command]; !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := commands[command].callback(cfg, param1)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays next page of location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous page of location areas in the Pokemon world",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Displays a list of Pokemon that reside in the given area",
			callback:    commandExplore,
		},
	}
}

func commandExit(*config, string) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp(*config, string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:ex")
	commands := getCommands()
	for _, command := range commands {
		fmt.Printf("%v: %v\n", command.name, command.description)
	}
	return nil
}

func commandMap(cfg *config, s string) error {
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

func commandMapb(cfg *config, s string) error {
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

func commandExplore(cfg *config, area string) error {
	if area == "" {
		return fmt.Errorf("Please enter a valid area")
	}
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%v/", area)
	locationArea, err := pokedata.GetLocationArea(url, cfg.Cache)
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
