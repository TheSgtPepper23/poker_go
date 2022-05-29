package main

import "fmt"

func main() {
	table := createTable()

	player1 := createPlayer("Player 1", 50)
	player2 := createPlayer("Player 2", 50)

	table.addPlayer(player1)
	table.addPlayer(player2)

	table.deal(2)

	table.dealRiverCards(3)

	table.acceptBet(20, 1)
	table.acceptBet(20, 0)

	deck := Deck{
		Card{face: "H", value: "A"},
		Card{face: "S", value: "2"},
		Card{face: "D", value: "4"},
		Card{face: "S", value: "7"},
		Card{face: "S", value: "3"},
		Card{face: "D", value: "5"},
		Card{face: "S", value: "9"},
	}

	fmt.Println(evaluateHand(deck, Deck{}))
}
