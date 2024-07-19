package main

import (
	"bufio"
	"fmt"
	"os"
)

type Command struct {
	name        string
	description string
	command     func()
}

func getCommands() map[string]Command {
	commandMap := map[string]Command{
		"exit": {
			name:        "exit",
			description: "Exit the program",
			command:     commandExit,
		},
		"help": {
			name:        "help",
			description: "Display help message",
			command:     commandHelp,
		},
	}
	return commandMap
}

func commandExit() {
	os.Exit(0)
}

func commandHelp() {
	scanner := bufio.NewScanner(os.Stdin)
	commandMap := getCommands()
	fmt.Println("")
	for _, command := range commandMap {
		fmt.Println(command.name + ": " + command.description)
	}
	fmt.Print("\nPress ENTER to continue > ")
	scanner.Scan()
}
