package main

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"pokedex": {
			name:        "pokedex",
			description: "Displays all the pokemon you caught",
			callback:    callBackPokedex,
		},
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
		"explore": {
			name:        "explore {location_area}",
			description: "Lists pokemon in a location",
			callback:    callBackExplore,
		},
		"catch": {
			name:        "catch {pokemon_name}",
			description: "Catch pokemon & add to your pokedex",
			callback:    callBackCatch,
		},
		"inspect": {
			name:        "inspect {pokemon_name}",
			description: "Inspect the pokemon stored in your pokedex",
			callback:    callBackInspect,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		
	}
}