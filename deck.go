package main

import "fmt"

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(card, i)
	}

}
