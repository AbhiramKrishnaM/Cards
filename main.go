package main

import "fmt"

func main() {
	cards := newdeck()

	// cards.print()
	// fmt.Println(cards)

	fmt.Println("Play hand")

	hand, remainingDeck := deal(cards, 2)
	fmt.Println(hand, " balance ", remainingDeck)

}
