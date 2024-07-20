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

func (d Duel) trapTriggersMod() int {
	if d.trapTriggers < 0 {
		fmt.Println("Trap Triggers are somehow negative, fix!!")
		return 0
	}
	if d.trapTriggers < 1 {
		return 2
	}
	if d.trapTriggers < 3 {
		return -8
	}
	if d.trapTriggers < 5 {
		return -16
	}
	if d.trapTriggers < 7 {
		return -24
	}
	return -32
}

func (d Duel) cardsUsedMod() int {
	if d.cardsUsed < 0 {
		fmt.Println("Cards Used are somehow negative, fix!!")
		return 0
	}
	if d.cardsUsed < 9 {
		return 15
	}
	if d.cardsUsed < 13 {
		return 12
	}
	if d.cardsUsed < 33 {
		return 0
	}
	if d.cardsUsed < 37 {
		return -5
	}
	return -7
}

func (d Duel) lpRemainingMod() int {
	if d.lpRemaining >= 8000 {
		return 6
	}
	if d.lpRemaining >= 7000 {
		return 4
	}
	if d.lpRemaining >= 1000 {
		return 0
	}
	if d.lpRemaining >= 100 {
		return -5
	}
	return -7
}

func (d Duel) getScore() int {
	return 50 + d.turnMod() + d.effectiveAttacksMod() + d.defensiveWinsMod() + d.facedownPlaysMod() + d.fusionsMod() + d.equipsMod() + d.magicsMod() + d.trapTriggersMod() + d.cardsUsedMod() + d.lpRemainingMod()
}

func (d Duel) getRank() string {
	score := d.getScore()
	if score <= 9 {
		return "S Tec"
	}
	if score <= 19 {
		return "A Tec"
	}
	if score <= 29 {
		return "B Tec"
	}
	if score <= 39 {
		return "C Tec"
	}
	if score <= 49 {
		return "D Tec"
	}
	if score <= 59 {
		return "D Pow"
	}
	if score <= 69 {
		return "C Pow"
	}
	if score <= 79 {
		return "B Pow"
	}
	if score <= 89 {
		return "A Pow"
	}
	return "S Pow"
}
