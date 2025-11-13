package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/sebmaz93/gokedex/internal/pokeapi"
)

type Config struct {
	ApiClient   pokeapi.Client
	NextUrl     *string
	PreviousUrl *string
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	return strings.Fields(lower)
}

func StartRepl(config *Config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			userInput := scanner.Text()
			words := cleanInput(userInput)
			if len(words) == 0 {
				fmt.Println("Your command is empty.")
				continue
			}
			cmd, err := GetCommand(words[0])
			if err != nil {
				fmt.Println(err)
				continue
			}
			err = cmd.Callback(config)
			if err != nil {
				fmt.Println(err)
			}
			continue
		}
	}
}
