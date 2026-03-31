package main

import "strings"

func cleanInput(text string) []string {
	words := strings.Split(strings.ToLower(text), " ")
	return words
}
