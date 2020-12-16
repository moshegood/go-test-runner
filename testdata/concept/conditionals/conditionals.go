package conditionals

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	val := 0
	switch card {
	case "ace":
		val = 11
	case "ten", "jack", "queen", "king":
		val = 10
	case "one":
		val = 1
	case "two":
		val = 2
	case "three":
		val = 3
	case "four":
		val = 4
	case "five":
		val = 5
	case "six":
		val = 6
	case "seven":
		val = 7
	case "eight":
		val = 8
	case "nine":
		val = 9
	}
	return val
}

// IsBlackjack returns true if the player has a blackjack, false otherwise.
func IsBlackjack(card1, card2 string) bool {
	panic("Please implement the IsBlackjack function")
}

// LargeHand implements the decision tree for hand scores larger than 20 points.
func LargeHand(isBlackjack bool, dealerScore int) string {
	panic("Please implement the LargeHand function")
}

// SmallHand implements the decision tree for hand scores with less than 21 points.
func SmallHand(handScore int, dealerScore int) string {
	panic("Please implement the SmallHand function")
}

// FirstTurn returns the semi-optimal decision for the first turn, given the cards of the player and the dealer.
// This function is already implemented and does not need to be edited. It pulls the other functions together in a
// complete decision tree for the first turn.
func FirstTurn(card1, card2, dealerCard string) string {
	handScore := ParseCard(card1) + ParseCard(card2)
	dealerScore := ParseCard(dealerCard)

	if 20 < handScore {
		return LargeHand(IsBlackjack(card1, card2), dealerScore)
	}
	return SmallHand(handScore, dealerScore)
}
