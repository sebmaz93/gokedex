package commands

import (
	"fmt"
	"os"
)

func commandMap() error {
	_, err := fmt.Printf("Poke maps")
	if err != nil {
		os.Exit(1)
		return err
	}
	return nil
}
