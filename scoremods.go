package main

import (
	"fmt"
)

func (d Duel) turnMod() int {
	if d.turns < 0 {
		fmt.Println("Turns are somehow negative, fix!!")
		return 0
	}
	if d.turns < 5 {
		return 12
	}
	if d.turns < 9 {
		return 8
	}
	if d.turns < 20 {
		return 0
	}
	if d.turns < 33 {
		return -8
	}
	return -12
}

func (d Duel) effectiveAttacksMod() int {
	if d.effectiveAttacks < 0 {
		fmt.Println("Effective Attacks are somehow negative, fix!!")
		return 0
	}
	if d.turns < 2 {
		return 4
	}
	if d.turns < 4 {
		return 2
	}
	if d.turns < 10 {
		return 0
	}
	if d.turns < 20 {
		return -2
	}
	return -4
}

func (d Duel) defensiveWinsMod() int {
	if d.effectiveAttacks < 0 {
		fmt.Println("Defensive Wins are somehow negative, fix!!")
		return 0
	}
	if d.turns < 2 {
		return 0
	}
	if d.turns < 6 {
		return -10
	}
	if d.turns < 10 {
		return -20
	}
	if d.turns < 15 {
		return -30
	}
	return -40
}
