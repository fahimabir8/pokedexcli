package main

import (
	"errors"
	"fmt"
)


func callBackExplore(cfg *config, args ...string) error{
	if len(args) != 1 {
		return errors.New("No location provided")
	}
	locationAreaName := args[0]

	locationArea, err := cfg.pokeapiClient.GetLocationArea(locationAreaName)

	if err != nil {
		return err
	}
	
	fmt.Printf("Exploring %s...\n", locationArea.Name)
	for _, pokemon := range locationArea.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}

	return nil
}