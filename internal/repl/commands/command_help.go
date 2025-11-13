package commands

import (
	"fmt"
)

func commandHelp() error {
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
