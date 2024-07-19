package main

import (
	"bufio"
	"fmt"
	"os"
)

func startPrompt() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Current rank: ")
		fmt.Println("Duel Points: ")
		fmt.Print("Action > ")
		scanner.Scan()
		input := scanner.Text()
		fmt.Println("echo...", input)
	}
}
