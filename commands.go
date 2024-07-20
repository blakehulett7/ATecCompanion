package main

import (
	"fmt"
	"os"
)

type Command struct {
	name        string
	description string
	command     func(*Duel)
}

const cyan = "\033[36m"
const reset = "\033[0m"

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
		"turnstatus": {
			name:        "turnstatus",
			description: "Display how many turns have passed this duel",
			command:     commandTurnStatus,
		},
		"effectiveattacksstatus": {
			name:        "effectiveattacksstatus",
			description: "Display how many effective attacks you've had this duel",
			command:     commandEffectiveAttacksStatus,
		},
	}
	return commandMap
}

func commandExit(duel *Duel) {
	os.Exit(0)
}

func commandHelp(duel *Duel) {
	commandMap := getCommands()
	fmt.Println("")
	for _, command := range commandMap {
		fmt.Println(command.name + ": " + command.description)
	}
}

func commandTurnStatus(duel *Duel) {
	fmt.Println("\nTurns Passed:", duel.turns)
	fmt.Println("Score Impact:", duel.turnMod())
	fmt.Println("")
	lineMap := map[string]int{
		"0-4":   12,
		"5-8":   8,
		"9-28":  0,
		"29-32": -8,
		">33":   -12,
	}
	for key, value := range lineMap {
		if duel.turnMod() == lineMap[key] {
			fmt.Printf(string(cyan)+"    %v: %v\n"+string(reset), key, value)
			continue
		}
		fmt.Printf("    %v: %v\n", key, value)
	}
}

func commandEffectiveAttacksStatus(duel *Duel) {
	fmt.Println("\nEffective Attacks", duel.effectiveAttacks)
	fmt.Println("Score Impact:", duel.effectiveAttacksMod())
	fmt.Println("")
	lineMap := map[string]int{
		"0-1":   12,
		"5-8":   8,
		"9-28":  0,
		"29-32": -8,
		">33":   -12,
	}
	for key, value := range lineMap {
		if duel.turnMod() == lineMap[key] {
			fmt.Printf(string(cyan)+"    %v: %v\n"+string(reset), key, value)
			continue
		}
		fmt.Printf("    %v: %v\n", key, value)
	}
}
