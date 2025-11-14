package main

import (
	"time"

	"github.com/sebmaz93/gokedex/internal/pokeapi"
	"github.com/sebmaz93/gokedex/internal/repl"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 10*time.Minute)
	config := &repl.Config{
		ApiClient: pokeClient,
		Pokebank:  map[string]pokeapi.Pokemon{},
	}
	repl.StartRepl(config)
}
