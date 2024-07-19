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
	if d.effectiveAttacks < 2 {
		return 4
	}
	if d.effectiveAttacks < 4 {
		return 2
	}
	if d.effectiveAttacks < 10 {
		return 0
	}
	if d.effectiveAttacks < 20 {
		return -2
	}
	return -4
}

func (d Duel) defensiveWinsMod() int {
	if d.defensiveWins < 0 {
		fmt.Println("Defensive Wins are somehow negative, fix!!")
		return 0
	}
	if d.defensiveWins < 2 {
		return 0
	}
	if d.defensiveWins < 6 {
		return -10
	}
	if d.defensiveWins < 10 {
		return -20
	}
	if d.defensiveWins < 15 {
		return -30
	}
	return -40
}

func (d Duel) facedownPlaysMod() int {
	if d.facedownPlays < 0 {
		fmt.Println("Facedown Plays are somehow negative, fix!!")
		return 0
	}
	if d.facedownPlays < 1 {
		return 0
	}
	if d.facedownPlays < 11 {
		return -2
	}
	if d.facedownPlays < 21 {
		return -4
	}
	if d.facedownPlays < 31 {
		return -6
	}
	return -8
}

func (d Duel) fusionsMod() int {
	if d.fusions < 0 {
		fmt.Println("Fusions are somehow negative, fix!!")
		return 0
	}
	if d.fusions < 1 {
		return 4
	}
	if d.fusions < 5 {
		return 0
	}
	if d.fusions < 10 {
		return -4
	}
	if d.fusions < 15 {
		return -8
	}
	return -12
}

func (d Duel) equipsMod() int {
	if d.equips < 0 {
		fmt.Println("Equips are somehow negative, fix!!")
		return 0
	}
	if d.equips < 1 {
		return 4
	}
	if d.equips < 5 {
		return 0
	}
	if d.equips < 10 {
		return -4
	}
	if d.equips < 15 {
		return -8
	}
	return -12
}

func (d Duel) magicsMod() int {
	if d.magics < 0 {
		fmt.Println("Magics are somehow negative, fix!!")
		return 0
	}
	if d.magics < 1 {
		return 2
	}
	if d.magics < 4 {
		return -4
	}
	if d.magics < 7 {
		return -8
	}
	if d.magics < 10 {
		return -12
	}
	return -16
}
