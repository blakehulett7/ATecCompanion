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
		score: 50,
	}
}

func getRank(score int) string {
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
