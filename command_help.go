package main

import "fmt"

func callbackHelp(_ *config) error {
	fmt.Println("Welcome to the Pokedex help menu!")
	fmt.Println("Here are the available commands")
	availableCommands := getCommands()
	for _, command := range availableCommands {
		fmt.Printf("%s - %s\n", command.name, command.description)
	}
	return nil
}
