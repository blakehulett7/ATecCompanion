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
			name:        "turn status",
			description: "Display how many turns have passed this duel",
			command:     commandTurnStatus,
		},
		"effective attack status": {
			name:        "effective attack status",
			description: "Display how many effective attacks you've had this duel",
			command:     commandEffectiveAttacksStatus,
		},
		"defensive wins status": {
			name:        "defensive wins status",
			description: "Display how many defensive wins you've had this duel",
			command:     commandDefensiveWinsStatus,
		},
		"facedown plays status": {
			name:        "facedown plays status",
			description: "Display how many facedown plays you've had this duel",
			command:     commandFacedownPlaysStatus,
		},
		"fusion status": {
			name:        "fusion status",
			description: "Display how many fusions you've had this duel",
			command:     commandFusionStatus,
		},
		"equip status": {
			name:        "equip status",
			description: "Display how many equips you've used this duel",
			command:     commandEquipStatus,
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

func commandDefensiveWinsStatus(duel *Duel) {
	fmt.Println("\nDefensive Wins", duel.defensiveWins)
	fmt.Println("Score Impact:", duel.defensiveWinsMod())
	fmt.Println("")
	factorLines := []FactorLine{
		{frequency: "0-1", scoremod: 0},
		{frequency: "2-5", scoremod: -10},
		{frequency: "6-9", scoremod: -20},
		{frequency: "10-14", scoremod: -30},
		{frequency: "over 15", scoremod: -40},
	}
	logRankFactor(duel, duel.defensiveWinsMod(), factorLines)
}

func commandFacedownPlaysStatus(duel *Duel) {
	fmt.Println("\nFacedown Plays", duel.facedownPlays)
	fmt.Println("Score Impact:", duel.facedownPlaysMod())
	fmt.Println("")
	factorLines := []FactorLine{
		{frequency: "0", scoremod: 0},
		{frequency: "1-10", scoremod: -2},
		{frequency: "11-20", scoremod: -4},
		{frequency: "21-30", scoremod: -6},
		{frequency: "over 30", scoremod: -8},
	}
	logRankFactor(duel, duel.facedownPlaysMod(), factorLines)
}

func commandFusionStatus(duel *Duel) {
	fmt.Println("\nFusions Played", duel.fusions)
	fmt.Println("Score Impact:", duel.fusionsMod())
	fmt.Println("")
	factorLines := []FactorLine{
		{frequency: "0", scoremod: 4},
		{frequency: "1-4", scoremod: 0},
		{frequency: "5-9", scoremod: -4},
		{frequency: "10-14", scoremod: -8},
		{frequency: "over 14", scoremod: -12},
	}
	logRankFactor(duel, duel.fusionsMod(), factorLines)
}

func commandEquipStatus(duel *Duel) {
	fmt.Println("\nEquips Played", duel.equips)
	fmt.Println("Score Impact:", duel.equipsMod())
	fmt.Println("")
	factorLines := []FactorLine{
		{frequency: "0", scoremod: 4},
		{frequency: "1-4", scoremod: 0},
		{frequency: "5-9", scoremod: -4},
		{frequency: "10-14", scoremod: -8},
		{frequency: "over 14", scoremod: -12},
	}
	logRankFactor(duel, duel.equipsMod(), factorLines)
}
