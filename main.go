package main

import "fmt"

func main() {
	cards := newdeck()

	fmt.Println(cards.saveToFile("new"))

	// newcard := newDeckFromFile("new")

	// newcard.print()

	cards.shuffle()
	cards.print()

	// cards.print()
	// fmt.Println(cards)

	// fmt.Println("Play hand")

	// hand, remainingDeck := deal(cards, 2)
	// fmt.Println(hand, " balance ", remainingDeck)

	// greeting := []byte("Hello World")

	// fmt.Println(greeting)

}
