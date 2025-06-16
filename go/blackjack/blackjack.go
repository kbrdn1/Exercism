package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
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
	case "ten", "jack", "queen", "king":
		return 10
	default:
	return 0
	}
}

// Deal returns the respective values of the two cards in a blackjack hand.
func Deal(card1, card2 string) int {
	card1Value := ParseCard(card1)
	card2Value := ParseCard(card2)
	return card1Value + card2Value
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	if card1 == "ace" && card2 == "ace" {
		return "P"
	}

	total := Deal(card1, card2)
	dealerCardValue := ParseCard(dealerCard)

	switch total {
		case 12:
			if dealerCardValue >= 4 && dealerCardValue <= 6 {
				return "S"
			}
			return "H"
		case 13, 14, 15, 16:
			if dealerCardValue >= 7 {
				return "H"
			}
			return "S"
		case 17:
			return "S"
		case 18:
			if dealerCardValue == 10 || dealerCardValue == 9 {
				return "H"
			}
			return "S"
		case 19, 20:
			return "S"
		case 21:
			if dealerCardValue == 10 || dealerCardValue == 11 {
				return "S"
			}
			return "W"
		default:
			return "H"
	}
}
