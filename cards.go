package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

//Represents a card with its face and value
type Card struct {
	face  string
	value string
}

//Returns a numeric value corresponding to the cards value
func (c Card) aValue() int {
	for i, v := range Values {
		if v == c.value {
			return i + 1
		}
	}

	return -1
}

//Represents a collection of cards
type Deck []Card

//Returns a new Deck containing the 52 cards of the game. It's not shuffled
func createDeck() Deck {
	cards := Deck{}
	for _, f := range Faces {
		for _, v := range Values {
			cards = append(cards, Card{value: v, face: f})
		}
	}

	return cards
}

//Randomizes the position of the cards inside the deck
func (d Deck) shuffle() {
	rand.Seed(time.Now().UnixNano())

	rand.Shuffle(len(d), func(j, i int) {
		d[i], d[j] = d[j], d[i]
	})
}

//Returns the top q cards of the deck and the rest of the cards
func deal(d Deck, q int) (Deck, Deck) {
	return d[:q], d[q:]
}

//Prints out the deck, its only for developing reasons
func (d Deck) print() {
	for i, c := range d {
		fmt.Println(i, fmt.Sprintln(c.face, c.value))
	}
}

//Sorts the cards on the deck by the value of the card. It ignores the face. The order parameter defines if the order is ascendant or descendant
func (d Deck) sort(order int) {
	if order >= 0 {
		sort.SliceStable(d, func(i, j int) bool {
			return d[i].aValue() < d[j].aValue()
		})
	} else {
		sort.SliceStable(d, func(i, j int) bool {
			return d[i].aValue() > d[j].aValue()
		})
	}
}

//Checks if the deck has any card with the value of the passed card
func (d Deck) hasValue(c Card) bool {
	for _, v := range d {
		if v.value == c.value {
			return true
		}
	}

	return false
}

// func (d Deck) hasFace(c Card) bool {
// 	for _, v := range d {
// 		if v.face == c.face {
// 			return true
// 		}
// 	}

// 	return false
// }

// func (d Deck) hasCard(c Card) bool {
// 	for _, v := range d {
// 		if v == c {
// 			return true
// 		}
// 	}

// 	return false
// }
