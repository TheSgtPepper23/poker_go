package main

import "github.com/lithammer/shortuuid/v3"

//Represent a player at the table with its hand and its chips
type Player struct {
	name  string
	hand  Deck
	chips int
	id    string
}

//Returns a new player with the specified name, the specified amount of chips, an empty hand and a random id
func createPlayer(name string, amount int) Player {
	p := Player{name: name, hand: make([]Card, 0), chips: amount, id: shortuuid.New()}
	return p
}

//The indicated amount is removed from the player chips and the same value is returned
func (p *Player) bet(amount int) int {
	p.chips = p.chips - amount

	return amount
}

//Repesents a table to play poker, it has the deck, can sit the players, has a river and a pot. It's supose to allow
//diferent poker games, so some elements may not be used in every ocation
type Table struct {
	deck    Deck
	players []Player
	river   Deck
	pot     int
}

//Returns a new table, with a shuffled deck, an empty list of players, and also empty river and pot
func createTable() Table {
	deck := createDeck()
	deck.shuffle()
	table := Table{
		deck:    deck,
		players: make([]Player, 0),
		river:   Deck{},
		pot:     0,
	}

	return table
}

//Sits a new player in the table if the current amount is less than 5 and returns a boolean indicating the result
//of the operation
func (t *Table) addPlayer(p Player) bool {
	if len(t.players) < 5 {
		t.players = append(t.players, p)
		return true
	} else {
		return false
	}
}

//Deals the indicated quantity of cards to each player in the table from the tables deck (The cards are removed from the deck)
func (t *Table) deal(q int) {
	for i := range t.players {
		hand, rest := deal(t.deck, q)
		t.deck = rest
		t.players[i].hand = hand
	}
}

//Deals the indicated amount of cards to the table river from the table deck (The cards are removed from the deck)
func (t *Table) dealRiverCards(q int) {
	if len(t.river) < 5 {
		river, rest := deal(t.deck, q)
		t.river = river
		t.deck = rest
	}
}

//Takes the bet from a player and adds that amount to the table pot
func (t *Table) acceptBet(amount, playerPosition int) {
	chips := t.players[playerPosition].bet(amount)
	t.pot += chips
}