package main

func main() {
	cards := newDeck()

	handed, remainingCards := deal(cards, 3)

	handed.print()
	remainingCards.print()
}
