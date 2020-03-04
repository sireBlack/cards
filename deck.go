package main

import "fmt"

// Create a new type of 'deck'
//which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "King", "Queen", "Joker"}

	for _, suite := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suite)
		}
	}

	return cards

}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}
