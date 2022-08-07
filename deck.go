package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type deck []string

func newdeck() deck {

	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}

	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Printing values of deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(card, i)
	}

}

// converting deck to string

func (d deck) convertToString() string {

	stringCard := strings.Join([]string(d), ",")

	return stringCard
}

func deal(d deck, handSize int) (deck, deck) {

	return d[:handSize], d[handSize:]

}

func (d deck) saveToFile(filename string) error {

	return ioutil.WriteFile(filename, []byte(d.convertToString()), 0666)

}
