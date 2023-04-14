package main

func main() {
	cards := deck{"Five of Heart", pickCard()}
	cards = append(cards, "Three of Spades")

	cards.print()
}

func pickCard() string {
	return "Ace of Spades"
}
