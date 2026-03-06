package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func callBackCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("No pokemon name provided")
	}
	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)

	if err != nil {
		return err
	}

	const threshold = 50

	randNum := rand.IntN(pokemon.BaseExperience)
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	if randNum > threshold {
		return fmt.Errorf("%s escaped!", pokemonName)
	}

	cfg.caughtPokemon[pokemonName] = pokemon
	fmt.Printf("%s was caught!\n", pokemonName)
	fmt.Printf("You may now inspect it with the inspect command.\n")
	return nil
}
