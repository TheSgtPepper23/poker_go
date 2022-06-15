package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

//Represents a card with its face and value
type Card struct {
	face     string
	value    string
	numValue int
}

//Represents a collection of cards
type Deck []Card

//Returns a new Deck containing the 52 cards of the game. It's not shuffled
func createDeck() Deck {
	cards := Deck{}
	for _, f := range Faces {
		for i, v := range Values {
			cards = append(cards, Card{value: v, face: f, numValue: i + 1})
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

func (d Deck) highestIndex() int {
	v := 0
	di := -1

	for i, c := range d {
		if c.numValue > v {
			di = i
			v = c.numValue
		}
	}

	return di
}

//Returns the top q cards of the deck and the rest of the cards
func deal(d Deck, q int) (Deck, Deck) {
	return d[:q], d[q:]
}

//Prints out the deck, its only for developing reasons
func (d Deck) print() {
	fmt.Println(d, "!")
}

//Sorts the cards on the deck by the value of the card. It ignores the face. The order parameter defines if the order is ascendant (1) or descendant (-1)
func (d Deck) sort(order int) {
	if order >= 0 {
		sort.SliceStable(d, func(i, j int) bool {
			return d[i].numValue < d[j].numValue
		})
	} else {
		sort.SliceStable(d, func(i, j int) bool {
			return d[i].numValue > d[j].numValue
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

func (d Deck) hasFace(c Card) bool {
	for _, v := range d {
		if v.face == c.face {
			return true
		}
	}

	return false
}

func (d Deck) hasCard(c Card) bool {
	for _, v := range d {
		if v == c {
			return true
		}
	}

	return false
}
