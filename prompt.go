package main

import (
	"bufio"
	"fmt"
	"os"
)

func startPrompt() {
	duelTracker := startDuelTracker()
	commandMap := getCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\nCurrent rank:", getRank(duelTracker.score))
		fmt.Println("Duel Points:", duelTracker.score)
		fmt.Print("\nAction > ")
		scanner.Scan()
		input := scanner.Text()
		if !inputValidator(input) {
			fmt.Println("invalid action, type 'help' for valid actions")
			continue
		}
		command := commandMap[input]
		command.command()
	}
}

func inputValidator(input string) bool {
	commandMap := getCommands()
	for command := range commandMap {
		if command == input {
			return true
		}
	}
	return false
}
