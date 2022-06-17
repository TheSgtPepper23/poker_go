package pepperPoker

import (
	"sort"
)

var Faces = []string{"S", "H", "D", "C"}
var Values = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
var Hands = []string{"highCard", "onePair", "twoPairs", "threeKind", "straight", "flush", "fullhouse", "fourKind", "straightFlush", "royalFlush"}

type Evaluation struct {
	face       string
	sum        int
	handRating int
}

func evaluateHand(p, river Deck) Evaluation {
	hand := append(p, river...)
	hand.sort(-1)
	isPair, pairValues := hasPair(hand)
	isThree, highest3 := hasThreeOrFour(hand, 3)
	isStraight, straightDeck := hasStraight(hand)
	isFlush := hasFlush(hand)
	isFour, highest4 := hasThreeOrFour(hand, 4)
	isSF, SFDeck, sfFace := hasStarightFlush(hand)

	if isSF && sumValues(SFDeck, 0, len(SFDeck)) == 55 {
		return Evaluation{face: sfFace, sum: 55, handRating: 9}
	}

	if isSF {
		return Evaluation{face: sfFace, sum: sumValues(SFDeck, 0, len(SFDeck)), handRating: 8}
	}

	if isFour {
		return Evaluation{handRating: 7, sum: highest4 * 4}
	}

	if isThree && isPair {
		return Evaluation{handRating: 6, sum: (highest3 * 3) + (pairValues[0] * 2)}
	}

	if isFlush {
		return Evaluation{handRating: 5, sum: sumValues(p, 0, len(p))}
	}

	if isStraight {
		return Evaluation{handRating: 4, sum: sumValues(straightDeck, 0, len(straightDeck))}
	}

	if isThree {
		return Evaluation{handRating: 3, sum: highest3 * 3}
	}

	if isPair && len(pairValues) > 1 {
		return Evaluation{handRating: 2, sum: pairValues[0]*2 + pairValues[1]*2}
	}

	if isPair {
		return Evaluation{handRating: 1, sum: pairValues[0] * 2}
	}

	return Evaluation{handRating: 0}
}

//Returns a map with the repetitions for each of the values present in the deck
func separateByValue(d Deck) map[int]int {
	m := make(map[int]int)

	for _, c := range d {
		q, v := m[c.numValue]
		if v {
			m[c.numValue] = q + 1
		} else {
			m[c.numValue] = 1
		}
	}

	return m
}

//Returns a map with the repetitions for each of the faces present in the deck
func separateByFace(d Deck) map[string]int {
	m := make(map[string]int)

	for _, c := range d {
		q, v := m[c.face]
		if v {
			m[c.face] = q + 1
		} else {
			m[c.face] = 1
		}
	}

	return m
}

//Determines how many pairs of cards of the same face the deck has, it also returns the value of every pair
func hasPair(hand Deck) (bool, []int) {

	values := []int{}
	has := false

	m := separateByValue(hand)

	for k, v := range m {
		if v == 2 {
			has = true
			values = append(values, k)
		}

	}

	//Descending sort of the values
	sort.Sort(sort.Reverse(sort.IntSlice(values)))

	return has, values
}

//Determines if the deck has q cards of the same value and the value of those cards
func hasThreeOrFour(hand Deck, q int) (bool, int) {

	higher := 0
	has := false

	m := separateByValue(hand)

	for k, v := range m {
		if v == q && k > higher {
			has = true
			higher = k
		}

	}

	return has, higher
}

//Returns the sum value of all the cards in the specified range for deck
func sumValues(d Deck, s, e int) int {
	sum := 0
	for _, v := range d[s:e] {
		sum += v.numValue
	}

	return sum
}

//Checks if the deck has a straigt and returns the sum of the values from that straight
func hasStraight(hand Deck) (bool, Deck) {
	hand = removeDuplicateValues(hand)
	hand.sort(-1)
	consecutiveCount, last, straigtStart := 0, 0, 0
	for i, c := range hand {
		if consecutiveCount == 4 {
			break
		}
		//If theres is a difference of 1 beetween the last and current value
		if last-c.numValue == 1 {
			if consecutiveCount == 0 {
				//Never can lower than 1, because the first comparison allwais will fail (last is 0 by default)
				straigtStart = i - 1
			}
			consecutiveCount++
		} else {
			consecutiveCount = 0
		}

		last = c.numValue
	}

	if consecutiveCount == 4 {
		return true, hand[straigtStart : straigtStart+5]
	} else {
		return false, Deck{}
	}
}

func hasStarightFlush(hand Deck) (bool, Deck, string) {
	hand.sort(-1)
	consecutiveCount, last, straigtStart := 0, 0, 0
	lastFace := ""
	for i, c := range hand {
		if consecutiveCount == 4 {
			break
		}
		//If theres is a difference of 1 beetween the last and current value
		if last-c.numValue == 1 && lastFace == c.face {
			if consecutiveCount == 0 {
				//Never can lower than 1, because the first comparison allwais will fail (last is 0 by default)
				straigtStart = i - 1
			}
			consecutiveCount++
		} else {
			consecutiveCount = 0
		}

		last = c.numValue
		lastFace = c.face
	}

	if consecutiveCount == 4 {
		return true, hand[straigtStart : straigtStart+5], lastFace
	} else {
		return false, Deck{}, lastFace
	}
}

//Removes the cards with the same value and only keeps one of each
func removeDuplicateValues(d Deck) Deck {
	nd := Deck{}

	for _, v := range d {
		if !nd.hasValue(v) {
			nd = append(nd, v)
		}
	}

	return nd
}

//Evaluates if the Deck has a flush and returns the deck if its necesary to compare two flushes
func hasFlush(d Deck) bool {
	d.sort(-1)
	m := separateByFace(d)

	for _, v := range m {
		if v >= 5 {
			return true
		}
	}

	return false
}

//Of a list of players (all winners) return the one with the highest card. If two have the same hand, then returns the first
//TODO Implement first round of untie in which the winner is decided by the highest value of his hand and not a kicker
func untie(players []Player, evaluations []Evaluation) int {
	defSum := evaluations[0].sum
	allEqual := true

	for _, e := range evaluations {
		if e.sum != defSum {
			allEqual = false
			break
		}
	}

	if allEqual {
		decks := Deck{}

		for _, p := range players {
			tempDeck := make(Deck, len(p.hand))
			copy(tempDeck, p.hand)
			decks = append(decks, tempDeck[tempDeck.highestIndex()])
		}

		return decks.highestIndex()
	} else {
		hs := 0
		hsi := 0

		for i, e := range evaluations {
			if e.sum > hs {
				hs = e.sum
				hsi = i
			}
		}

		return hsi
	}

}
