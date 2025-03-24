package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	// The goal of this exercise is to applying the `switch... case` concept
	// But I'd use a Map structure for this particular problem.
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten":
		return 10
	case "jack":
		return 10
	case "queen":
		return 10
	case "king":
		return 10
	default:
		return 0
	}
}

func isInRange(start, end, value int) bool {
	return start <= value && value <= end
}

func isBlackJack(cardsSum int) bool {
	return cardsSum == 21
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	cardsSum := ParseCard(card1) + ParseCard(card2)

	switch {
	case card1 == "ace" && card2 == "ace":
		return "P"
	case isBlackJack(cardsSum):
		if ParseCard(dealerCard) < 10 {
			return "W"
		} else {
			return "S"
		}
	case isInRange(17, 20, cardsSum):
		return "S"
	case isInRange(12, 16, cardsSum):
		if ParseCard(dealerCard) >= 7 {
			return "H"
		} else {
			return "S"
		}
	default:
		return "H"
	}
}
