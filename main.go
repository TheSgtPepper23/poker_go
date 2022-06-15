package main

import "fmt"

func main() {
	table := createTable(5)

	player1 := createPlayer("Player 1", 50)
	player2 := createPlayer("Player 2", 50)
	player3 := createPlayer("Player 3", 50)
	player4 := createPlayer("Player 4", 50)
	player5 := createPlayer("Player 5", 50)

	table.addPlayer(player1)
	table.addPlayer(player2)
	table.addPlayer(player3)
	table.addPlayer(player4)
	table.addPlayer(player5)

	table.deal(2)

	table.dealRiverCards(5)

	fmt.Println(table.showTime())

}
