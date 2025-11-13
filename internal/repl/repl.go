package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	cmds "github.com/sebmaz93/gokedex/internal/repl/commands"
)

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	return strings.Fields(lower)
}

func StartRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	config := cmds.Config{}
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			userInput := scanner.Text()
			words := cleanInput(userInput)
			if len(words) == 0 {
				fmt.Println("Your command is empty.")
				continue
			}
			cmd, err := cmds.GetCommand(words[0])
			if err != nil {
				fmt.Println(err)
				continue
			}
			err = cmd.Callback(&config)
			if err != nil {
				fmt.Println(err)
			}
			continue
		}
	}
}
