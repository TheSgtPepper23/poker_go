package pepperPoker

import (
	"testing"
)

func TestDeck(t *testing.T) {
	newDeck := createDeck()

	if newDeck[0].face != "S" || newDeck[0].numValue != 1 {
		t.Errorf(newDeck[0].face)
	}

	newDeck.sort(-1)

	if newDeck[0].face == newDeck[1].face && newDeck[2].face == newDeck[1].face && newDeck[3].face == newDeck[2].face {
		t.Errorf("Is not properly shuffled")
	}

	newDeck.sort(-1)

	if newDeck[0].numValue != 13 {
		t.Errorf("dfsa")
	}

	newDeck.sort(1)

	if newDeck[0].numValue != 1 {
		t.Errorf("dfsa")
	}
}
