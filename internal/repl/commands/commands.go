package commands

import "fmt"

type cliCommand struct {
	Name        string
	Description string
	Callback    func(*Config) error
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"map": {
			Name:        "map",
			Description: "List map areas",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "List previous map areas",
			Callback:    commandMapB,
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

type Config struct {
	Next     string
	Previous string
}
