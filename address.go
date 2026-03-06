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

func startAddress(cfg *config) {
	// scanner := bufio.NewScanner(os.Stdin)
	// commands := getCommands()

	// for {

	// 	fmt.Print("Pokedex > ")
	// 	if scanner.Scan() {
	// 		res := cleanInput(scanner.Text())
	// 		if len(res) == 0 {
	// 			continue
	// 		}
	// 		commandName := res[0]

	// 		args := []string{}

	// 		if len(res) > 1 {
	// 			args = res[1:]
	// 		}

	// 		if len(res) == 1 {
	// 			if cmd, exists := commands[commandName]; exists {
	// 				err := cmd.callback(cfg, args...)
	// 				if err != nil {
	// 					fmt.Println(err)
	// 				}
	// 			} else {
	// 				fmt.Println("Unknown command")
	// 			}
	// 		}

	// 	}
	// }

	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}


