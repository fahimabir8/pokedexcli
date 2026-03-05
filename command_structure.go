package main

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    getHelp,
		},
		"map": {
			name:        "map",
			description: "List Next Location Areas",
			callback:    callBackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List Previous Location Areas",
			callback:    callBackMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		
	}
}