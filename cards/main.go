package main

func main() {

	cards := newDeck()

	hand, remainingCard := deal(cards, 5)
	// deal returns two values and the 1st will be assigned to hand
	// 2nd value will be assigned to remainingCard
	hand.print()
	remainingCard.print()
}

func newCard() string {
	return "Ace of spades"
}
