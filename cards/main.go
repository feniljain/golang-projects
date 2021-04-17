package main

func main() {
	//cards := newDeck()
	//cards = append(cards, "Six of spades")

	//hand, remainingCards := deal(cards, 5)
	//hand.print()
	//remainingCards.print()
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
