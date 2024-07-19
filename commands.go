package main

import "os"

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
	}
	return commandMap
}

func commandExit() {
	os.Exit(0)
}
