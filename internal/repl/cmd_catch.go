package repl

import (
	"fmt"
	"math/rand"
)

func commandCatch(config *Config, param string) error {
	if param == "" {
		return fmt.Errorf("Pokemon name missing")
	}

	pokemon, err := config.ApiClient.GetPokemon(param)
	if err != nil {
		return err
	}
	status := "escaped!"
	i := rand.Intn(pokemon.BaseExperience)
	if i < (pokemon.BaseExperience / 2) {
		status = "was caught!"
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	fmt.Printf("%s %s\n", pokemon.Name, status)
	fmt.Println("")

	config.Pokebank[pokemon.Name] = pokemon

	return nil
}
