package main

import (
	"bufio"
	"fmt"
	"strings"

	// "log"
	"os"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex > ")

		scanner.Scan()

		text := scanner.Text()

		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("Invalid Command")
			continue
		}

		command.callback()
	}
}

type cliCommand struct {
	name string
	description string
	callback func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand {
		"help": {
			name: "help",
			description: "Prints the help menu",
			callback: commandHelp,
		},
		"exit": {
			name: "exit",
			description: "Exits the Pokedex",
			callback: commandExit,
		},
	}
}

func cleanInput(str string) []string {
	loweredStr := strings.ToLower(str)
	words := strings.Fields(loweredStr)
	return words
}