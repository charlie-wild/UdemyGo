package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

/* Create a new type of deck which is a slice of strings */

type deck []string

func newDeck() deck {
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

func (d deck) print() { // any variable of type deck gets access to this print method using the receiver
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")

}

func (d deck) saveToFile(filename string) error {
	byteDeck := []byte(d.toString())
	return ioutil.WriteFile(filename, byteDeck, 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// log error and return a call to newDeck()
		// or log error and quit
		fmt.Println("Error:", err)
		os.Exit(1)

	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}
