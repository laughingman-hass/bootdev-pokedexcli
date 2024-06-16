package main

import (
	"errors"
	"fmt"
)

func callbackMap(cfg *config, _ ...string) error {
	locationAreas, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}

	fmt.Println("Location areas:")
	for _, area := range locationAreas.Results {
		fmt.Printf("- %s\n", area.Name)
	}

	cfg.nextLocationAreaURL = locationAreas.Next
	cfg.prevLocationAreaURL = locationAreas.Previous

	return nil
}

func callbackMapb(cfg *config, _ ...string) error {
	if cfg.prevLocationAreaURL == nil {
		return errors.New("previous URL not set")
	}

	locationAreas, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationAreaURL)
	if err != nil {
		return err
	}

	fmt.Println("Location areas:")
	for _, area := range locationAreas.Results {
		fmt.Printf("- %s\n", area.Name)
	}

	cfg.nextLocationAreaURL = locationAreas.Next
	cfg.prevLocationAreaURL = locationAreas.Previous

	return nil
}
