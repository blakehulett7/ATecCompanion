package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func startPrompt() {
	duelTracker := startDuelTracker()
	commandMap := getCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Welcome to the A Tec Companion! Use this to get quick A Tecs and great cards!")
		fmt.Println("\nCurrent rank:", duelTracker.getRank())
		fmt.Println("Duel Points:", duelTracker.getScore())
		fmt.Print("\nAction > ")
		scanner.Scan()
		input := scanner.Text()
		if !inputValidator(input) {
			fmt.Println("\ninvalid action, type 'help' for valid actions")
			fmt.Print("\nPress ENTER to continue > ")
			scanner.Scan()
			clear := exec.Command("clear")
			clear.Stdout = os.Stdout
			clear.Run()
			continue
		}
		command := commandMap[input]
		command.command(&duelTracker)
		fmt.Print("\nPress ENTER to continue > ")
		scanner.Scan()
		clear := exec.Command("clear")
		clear.Stdout = os.Stdout
		clear.Run()
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
