package main

func main() {
	cfg := config{
		Next:     nil,
		Previous: nil,
	}
	startRepl(&cfg)
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	Next     *string
	Previous *string
}
