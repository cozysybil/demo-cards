package main

import "fmt"

func main() {
	cards := []string{"Five of Heart", pickCard()}
	cards = append(cards, "Three of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func pickCard() string {
	return "Ace of Spades"
}
