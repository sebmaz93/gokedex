package repl

import "fmt"

type cliCommand struct {
	Name        string
	Description string
	Callback    func(*Config, string) error
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"map": {
			Name:        "map",
			Description: "List location areas",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "List previous location areas",
			Callback:    commandMapB,
		}, "explore": {
			Name:        "explore",
			Description: "Explore location area by {location_name}",
			Callback:    commandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "Catch a Pokemon by {pokemon_name}",
			Callback:    commandCatch,
		},
		"help": {
			Name:        "help",
			Description: "Display a help message",
			Callback:    commandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the pokedex",
			Callback:    commandExit,
		},
	}
}

func GetCommand(name string) (cliCommand, error) {
	cmd, ok := GetCommands()[name]
	if !ok {
		return cliCommand{}, fmt.Errorf("Unknown command")
	}
	return cmd, nil
}
