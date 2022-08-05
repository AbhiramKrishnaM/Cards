package main

import "fmt"

func main() {
	cards := []string{newCard(), newCard(), newCard()}

	cards = append(cards, "Six of spades")

	for i, card := range cards {
		fmt.Println(card, i)
	}

	fmt.Println(cards)

}

func newCard() string {
	return "Ace of Spades"
}
