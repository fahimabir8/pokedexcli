package main

import (
	"errors"
	"fmt"
)

func callBackInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("No pokemon name provided")
	}
	pokemonName := args[0]

	pokemon, ok := cfg.caughtPokemon[pokemonName]

	if !ok {
		return fmt.Errorf("you have not caught the pokemon")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _,pk := range pokemon.Stats {
		
			fmt.Printf(" -%v: %v\n",pk.Stat.Name,pk.BaseStat)
	}
	
	fmt.Println("Types:")
	for _,pk := range pokemon.Types {
		fmt.Printf("\t- %s\n",pk.Type.Name)
	}

	return nil
}