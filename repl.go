package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/thetramp22/pokedexcli/internal/pokeData"
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
		commands := getCommands()
		if _, ok := commands[command]; !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := commands[command].callback(cfg)
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
	}
}

func commandExit(*config) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp(*config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:ex")
	commands := getCommands()
	for _, command := range commands {
		fmt.Printf("%v: %v\n", command.name, command.description)
	}
	return nil
}

func commandMap(cfg *config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.Next != nil {
		url = *cfg.Next
	}
	if cfg.Previous != nil && cfg.Next == nil {
		fmt.Println("You're on the last page")
		return nil
	}

	locationArea, err := pokeData.GetLocationArea(url)
	if err != nil {
		fmt.Println(err)
	}
	for _, result := range locationArea.Results {
		fmt.Println(result.Name)
	}
	cfg.Next = locationArea.Next
	cfg.Previous = locationArea.Previous
	return nil
}

func commandMapb(cfg *config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.Previous != nil {
		url = *cfg.Previous
	}
	if cfg.Previous == nil {
		fmt.Println("You're on the first page")
		return nil
	}

	locationArea, err := pokeData.GetLocationArea(url)
	if err != nil {
		fmt.Println(err)
	}
	for _, result := range locationArea.Results {
		fmt.Println(result.Name)
	}
	cfg.Next = locationArea.Next
	cfg.Previous = locationArea.Previous
	return nil
}
