package main

import (
	"fmt"
)

func callBackPokedex(cfg *config, args ...string) error {
	fmt.Println("Showing your pokedex...")

	for _, pk := range cfg.caughtPokemon {
		fmt.Printf(" -%s\n", pk.Name)
	}
	
	return nil
}