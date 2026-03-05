package main

import (
	"time"

	"github.com/fahimabir8/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
}

func main() {
	cfg := config {
		pokeapiClient: pokeapi.NewClient(time.Hour),
	}
	startAddress(&cfg)
}
