package repl

import (
	"fmt"
)

func commandHelp(config *Config, param1 string) error {
	fmt.Printf(`
Welcome to the Pokedex!
Usage:

`)
	for _, cmd := range GetCommands() {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
	}
	fmt.Println()
	return nil
}
