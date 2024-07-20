package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Duel struct {
	turns            int
	effectiveAttacks int
	defensiveWins    int
	facedownPlays    int
	fusions          int
	equips           int
	magics           int
	trapTriggers     int
	cardsUsed        int
	lpRemaining      int
}

func startDuelTracker() Duel {
	return Duel{
		turns:            0,
		effectiveAttacks: 0,
		defensiveWins:    0,
		facedownPlays:    0,
		fusions:          0,
		equips:           0,
		magics:           0,
		trapTriggers:     0,
		cardsUsed:        0,
		lpRemaining:      8000,
	}
}

func commandPlayCards(duel *Duel) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("\nHow many fusions did you make? > ")
	scanner.Scan()
	fusionInput := scanner.Text()
	fusions, _ := strconv.Atoi(fusionInput)
	duel.cardsUsed += fusions
	duel.fusions += fusions
	fmt.Print("\nHow many equips did you use? > ")
	scanner.Scan()
	equipInput := scanner.Text()
	equips, _ := strconv.Atoi(equipInput)
	duel.cardsUsed += equips
	duel.equips += equips
	fmt.Print("\nHow many cards were discarded or played vanilla? > ")
	scanner.Scan()
	cardInput := scanner.Text()
	cards, _ := strconv.Atoi(cardInput)
	duel.cardsUsed += cards
}

func commandSetCard(duel *Duel) {
	duel.cardsUsed += 1
	duel.facedownPlays += 1
}

func commandPlayMagic(duel *Duel) {
	duel.cardsUsed++
	duel.magics += 1
}

func commandEffectiveAttack(duel *Duel) {
	duel.effectiveAttacks++
}

func commandDefensiveWin(duel *Duel) {
	duel.defensiveWins++
}

func commandTriggerTrap(duel *Duel) {
	duel.trapTriggers++
}

func commandEndTurn(duel *Duel) {
	duel.turns++
}

func commandLoseLp(duel *Duel) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("\nHow many life points did you lose? > ")
	scanner.Scan()
	input := scanner.Text()
	lostLP, _ := strconv.Atoi(input)
	duel.lpRemaining -= lostLP
}
