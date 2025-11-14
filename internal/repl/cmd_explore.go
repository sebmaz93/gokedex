package repl

import "fmt"

func commandExplore(config *Config, param string) error {
	if param == "" {
		return fmt.Errorf("Area missing")
	}

	d, err := config.ApiClient.ExploreLocation(param)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", param)
	fmt.Println("Found Pokemon:")
	for _, p := range d.PokemonEncounters {
		fmt.Printf("- %v\n", p.Pokemon.Name)
	}
	fmt.Println("")

	return nil
}
