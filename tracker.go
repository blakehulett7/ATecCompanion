package main

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
