package main

import (
	"fmt"
	"os"
)

const cyan = "\033[36m"
const reset = "\033[0m"

type Command struct {
	name        string
	description string
	command     func(*Duel)
}

type FactorLine struct {
	frequency string
	scoremod  int
}

func logRankFactor(duel *Duel, currentScoreMod int, factorLines []FactorLine) {
	for _, line := range factorLines {
		if currentScoreMod == line.scoremod {
			fmt.Printf(string(cyan)+"  > %v: %v\n"+string(reset), line.frequency, line.scoremod)
			continue
		}
		fmt.Printf("    %v: %v\n", line.frequency, line.scoremod)
	}
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
		"turn status": {
			name:        "turnstatus",
			description: "Display how many turns have passed this duel",
			command:     commandTurnStatus,
		},
		"effective attack status": {
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
	factorLines := []FactorLine{
		{frequency: "0-4", scoremod: 12},
		{frequency: "5-8", scoremod: 8},
		{frequency: "9-28", scoremod: 0},
		{frequency: "29-32", scoremod: -8},
		{frequency: "over 33", scoremod: -12},
	}
	logRankFactor(duel, duel.turnMod(), factorLines)
}

func commandEffectiveAttacksStatus(duel *Duel) {
	fmt.Println("\nEffective Attacks", duel.effectiveAttacks)
	fmt.Println("Score Impact:", duel.effectiveAttacksMod())
	fmt.Println("")
	factorLines := []FactorLine{
		{frequency: "0-1", scoremod: 4},
		{frequency: "5-8", scoremod: 2},
		{frequency: "9-28", scoremod: 0},
		{frequency: "29-32", scoremod: -2},
		{frequency: "over 33", scoremod: -4},
	}
	logRankFactor(duel, duel.effectiveAttacksMod(), factorLines)
}
