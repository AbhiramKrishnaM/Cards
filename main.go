package main

import "fmt"

func main() {
	cards := deck{newCard(), newCard(), newCard()}

	cards = append(cards, "Six of spades")

	cards.print()
	fmt.Println(cards)

}

func newCard() string {
	return "Ace of Spades"
}
