package main

import (
	"fmt"
)

func callbackExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("no location area provided")
	}

	locationAreaName := args[0]

	resp, err := cfg.pokeapiClient.GetLocationAreas(locationAreaName)
	if err != nil {
		return err
	}

	fmt.Printf("Pokemon in %s:\n", resp.Name)
	for _, pokemon := range resp.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
