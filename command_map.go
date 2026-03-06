package main

import (
	"errors"
	"fmt"
)


func callBackMap(cfg *config, args ...string) error{
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)

	if err != nil {
		return err
	}
	
	
	for _, area := range resp.Results {
		fmt.Printf("- %s\n", area.Name)
	}

	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previoius

	return err
}

func callBackMapb(cfg *config, args ...string) error{
	if cfg.prevLocationAreaURL == nil {
		errors.New("You're literally on the first page buddy!!")
	}
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationAreaURL)

	if err != nil {
		return err
	}
	
	for _, area := range resp.Results {
		fmt.Printf("%s\n", area.Name)
	}

	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previoius

	return err
}

