package main

import (
	"bufio"
	"fmt"
	"os"
)

type Command struct {
	name        string
	description string
	command     func(*Duel)
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
		"status": {
			name:        "status",
			description: "Display current state of the duel",
			command:     commandStatus,
		},
	}
	return commandMap
}

func commandExit(duel *Duel) {
	os.Exit(0)
}

func commandHelp(duel *Duel) {
	scanner := bufio.NewScanner(os.Stdin)
	commandMap := getCommands()
	fmt.Println("")
	for _, command := range commandMap {
		fmt.Println(command.name + ": " + command.description)
	}
	fmt.Print("\nPress ENTER to continue > ")
	scanner.Scan()
}

func commandStatus(duel *Duel) {
	return
}
