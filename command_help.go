package main

import (
	"fmt"
)

func commandHelp() error {
	fmt.Println("Welcome to the help menu for this Pokedex")
	fmt.Println("Here are the available commands:")
	
	// availableCommands := getCommands()
	for _, command := range getCommands() {
		fmt.Printf(" - %s: %s\n", command.name, command.description)
	}
	return nil
}