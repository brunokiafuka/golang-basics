package main

func main() {

	//cards := newDeck()
	cards := newDeckFromFile("my_cards")

	/* hand, remainingCard := deal(cards, 5)
	// deal returns two values and the 1st will be assigned to hand
	// 2nd value will be assigned to remainingCard
	hand.print()
	remainingCard.print() */
	cards.shuffle()
	cards.print()
	// cards.saveToFile("my_cards")
}

func newCard() string {
	return "Ace of spades"
}
