package repl

import "fmt"

func commandPokedex(config *Config, params string) error {

	if len(config.Pokebank) == 0 {
		fmt.Println("Oops your Pokedex is empty.")
	}

	fmt.Println("Your Pokedex:")
	for _, v := range config.Pokebank {
		fmt.Printf("- %s", v.Name)
	}
	fmt.Println("")

	return nil
}
