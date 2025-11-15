package repl

import "fmt"

func commandInspect(config *Config, param string) error {
	if param == "" {
		return fmt.Errorf("Pokemon name missing.")
	}

	pokemon, ok := config.Pokebank[param]
	if !ok {
		fmt.Println("Oops you have never caught this pokemon.")
		return nil
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, s := range pokemon.Stats {
		fmt.Printf("-%s: %v\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, t := range pokemon.Types {
		fmt.Printf("- %v\n", t.Type.Name)
	}
	fmt.Println("")

	return nil
}
