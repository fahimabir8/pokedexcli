package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)



func cleanInput(text string) []string {

	words := strings.Fields(strings.ToLower(strings.TrimSpace(text)))

	return words
}

func startAddress() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()

	for {

		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			res := cleanInput(scanner.Text())
			commandName := res[0]
			if len(res) == 0 {
				continue
			}

			if len(res) == 1 {
				if cmd, exists := commands[commandName]; exists {
					cmd.callback()
				} else {
					fmt.Println("Unknown command")
				}
			}

		}
	}

}


