package main

func main() {
	cards := newDeck()
	cards.print()
}

func pickCard() string {
	return "Ace of Spades"
}
