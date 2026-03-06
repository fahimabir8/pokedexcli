package main

import (
	"fmt"
)

func getHelp(cfg *config, args ...string) error {
	commands := getCommands()
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")

	for _, cmd := range commands {

		fmt.Printf("%s: %s\n",cmd.name, cmd.description)
	}
	return nil
}
