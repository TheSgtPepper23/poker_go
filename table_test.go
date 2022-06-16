package main

import (
	"fmt"
	"testing"
)

func TestPlayer(t *testing.T) {
	p1 := createPlayer("player 1", 500)

	if p1.name != "player 1" || p1.chips != 500 || len(p1.hand) != 0 {
		t.Errorf("Error on creation")
	}

	ba := p1.bet(200)

	if ba != 200 || p1.chips != 300 {
		t.Errorf("Something wrong")
	}
}

func TestTable(t *testing.T) {

	cardsToDeal := 2
	riverSize := 5

	table := createTable(5)

	if table.deck == nil || len(table.players) != 0 {
		t.Errorf("Error on creation")
	}

	p1 := createPlayer("player 1", 500)
	p2 := createPlayer("player 2", 500)
	p3 := createPlayer("player 3", 500)
	p4 := createPlayer("player 4", 500)
	p5 := createPlayer("player 5", 500)
	p6 := createPlayer("player 6", 500)

	table.addPlayer(p1)
	table.addPlayer(p2)
	table.addPlayer(p3)
	table.addPlayer(p4)
	table.addPlayer(p5)
	table.addPlayer(p6)

	if len(table.players) > 5 {
		t.Errorf("Too many players")
	}

	table.deal(cardsToDeal)

	for _, p := range table.players {
		if len(p.hand) != cardsToDeal {
			t.Errorf("Doesnt have the right amount of cards")
			return
		}
	}

	table.dealRiverCards(riverSize)

	if len(table.river) != riverSize {
		t.Errorf("Didnt deal the river")
	}

	if len(table.deck) != 52-(len(table.players)*cardsToDeal)-riverSize {
		t.Errorf("Deck didnt reduce its size")
	}

	table.acceptBet(200, 1)

	if table.pot != 200 {
		t.Errorf("Pot wrong amount")
	}

	if table.players[1].chips != 300 {
		t.Errorf(("Playe chips didnt decrease"))
	}

	//TODO test showtime with a rigged river and hands for every player
	p, s := table.showTime()
	table.river.print()

	fmt.Println(p, s, "hoa")
}

func TestShowTimeFunc(t *testing.T) {
	table := createTable(5)

	p1 := createPlayer("p1", 500)
	p2 := createPlayer("p2", 500)
	p3 := createPlayer("p3", 500)
	p4 := createPlayer("p4", 500)
	p5 := createPlayer("p5", 500)

	table.addPlayer(p1)
	table.addPlayer(p2)
	table.addPlayer(p3)
	table.addPlayer(p4)
	table.addPlayer(p5)

	testMultipleWinners(&table)

	fmt.Println(table.showTime())
	table.river.print()

}

func testMultipleWinners(t *Table) {
	t.river = append(
		t.river,
		Card{face: "H", value: "K", numValue: 12},
		Card{face: "H", value: "7", numValue: 6},
		Card{face: "D", value: "7", numValue: 6},
		Card{face: "D", value: "Q", numValue: 11},
		Card{face: "C", value: "3", numValue: 2},
	)

	t.players[0].hand = append(
		t.players[0].hand,
		Card{face: "C", value: "10", numValue: 9},
		Card{face: "C", value: "J", numValue: 10},
	)

	t.players[1].hand = append(
		t.players[1].hand,
		Card{face: "C", value: "2", numValue: 1},
		Card{face: "S", value: "3", numValue: 2},
	)

	t.players[2].hand = append(
		t.players[2].hand,
		Card{face: "H", value: "2", numValue: 1},
		Card{face: "D", value: "A", numValue: 13},
	)

	t.players[3].hand = append(
		t.players[3].hand,
		Card{face: "D", value: "10", numValue: 9},
		Card{face: "S", value: "K", numValue: 12},
	)
	t.players[4].hand = append(
		t.players[4].hand,
		Card{face: "C", value: "A", numValue: 13},
		Card{face: "D", value: "J", numValue: 10},
	)
}
