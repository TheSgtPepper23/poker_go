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
		Card{face: "C", value: "A"},
		Card{face: "S", value: "K"},
		Card{face: "S", value: "Q"},
		Card{face: "S", value: "J"},
		Card{face: "S", value: "10"},
		Card{face: "S", value: "9"},
		Card{face: "S", value: "5"},
	}

	fmt.Println(evaluateHand(deck, Deck{}))
}
