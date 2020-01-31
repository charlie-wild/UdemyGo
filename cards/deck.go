package main

import "fmt"

/* Create a new type of deck which is a slice of strings */

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for i, suit := range cardSuits {
		for j, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
}

func (d deck) print() { // any variable of type deck gets access to this print method using the receiver
	for i, card := range d {
		fmt.Println(i, card)
	}
}
