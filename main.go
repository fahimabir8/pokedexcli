package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"help": {
		name: "help",
		description: "Get help on your Journey",
		callback: getHelp,
	},
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			res := cleanInput(scanner.Text())
			
			if len(res) == 1 {
				if cmd, exists := commands[res[0]]; exists {
					cmd.callback()
				} else {
					fmt.Println("Unknown command")
				}
			}
			
		}
	}

}

func commandExit() error{
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil	
}

func getHelp() error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex")
	return nil
}