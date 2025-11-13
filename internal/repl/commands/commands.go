package commands

import "fmt"

type cliCommand struct {
	Name        string
	Description string
	Callback    func() error
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the pokedex",
			Callback:    commandExit,
		},
		"help": {
			Name:        "help",
			Description: "Display a help message",
			Callback:    commandHelp,
		},
		"map": {
			Name:        "map",
			Description: "List region maps",
			Callback:    commandMap,
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
