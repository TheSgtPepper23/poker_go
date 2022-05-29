package main

var Faces = []string{"S", "H", "D", "C"}
var Values = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
var Hands = []string{"highCard", "onePair", "twoPairs", "threeKind", "straight", "flush", "fullhouse", "fourKind", "straightFlush", "royalFlush"}

func determineValue(val string) int {
	for i, v := range Values {
		if v == val {
			return i + 1
		}
	}

	return -1
}

func evaluateHand(hand, river Deck) (string, Card) {
	wholeCards := append(hand, river...)
	wholeCards.sort(-1)
	higher := wholeCards[0]

	isPair, pairValues := hasPair(wholeCards)
	isThree, highest3 := hasThreeOrFour(wholeCards, 3)
	isStraight, sumStraight := hasStraight(wholeCards)
	isFlush, flushedDeck := hasFlush(wholeCards)
	isFour, highest4 := hasThreeOrFour(wholeCards, 4)

	if isFlush && isStraight && sumStraight == 55 {
		return Hands[9], higher
	}

	//TODO is wrong because the flush and the straigh can be non related
	if isFlush && isStraight {
		return Hands[8], higher
	}

	if isFour {
		return Hands[7], Card{face: "S", value: Values[highest4-1]}
	}

	if isThree && isPair {
		return Hands[6], Card{}
	}

	if isFlush {
		return Hands[5], flushedDeck[0]
	}

	if isStraight {
		return Hands[4], flushedDeck[0]
	}

	if isThree {
		return Hands[3], Card{face: "S", value: Values[highest3-1]}
	}

	if isPair && len(pairValues) > 1 {
		return Hands[2], Card{}
	}

	if isPair {
		return Hands[1], Card{}
	}

	return Hands[0], higher
}

//Returns a map with the repetitions for each of the values present in the deck
func separateByValue(d Deck) map[string]int {
	m := make(map[string]int)

	for _, c := range d {
		q, v := m[c.value]
		if v {
			m[c.value] = q + 1
		} else {
			m[c.value] = 1
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
			values = append(values, determineValue(k))
		}

	}

	return has, values
}

//Determines if the deck has q cards of the same value and the value of those cards
func hasThreeOrFour(hand Deck, q int) (bool, int) {

	higher := 0
	has := false

	m := separateByValue(hand)

	for k, v := range m {
		if v == q && determineValue(k) > higher {
			has = true
			higher = determineValue(k)
		}

	}

	return has, higher
}

//Returns the sum value of all the cards in the specified range for deck
func sumValues(d Deck, s, e int) int {
	sum := 0
	for _, v := range d[s:e] {
		sum += v.aValue()
	}

	return sum
}

//Checks if the deck has a straigt and returns the sum of the values from that straight
func hasStraight(hand Deck) (bool, int) {

	hand = removeDuplicateValues(hand)
	hand.sort(-1)
	consecutiveCount, last, straigtStart := 0, 0, 0
	for i, c := range hand {
		if consecutiveCount == 4 {
			break
		}
		//If theres is a difference of 1 beetween the last and current value
		if last-c.aValue() == 1 {
			if consecutiveCount == 0 {
				//Never can lower than 1, because the first comparison allwais will fail (last is 0 by default)
				straigtStart = i - 1
			}
			consecutiveCount++
		} else {
			consecutiveCount = 0
		}

		last = c.aValue()
	}

	if consecutiveCount == 4 {
		return true, sumValues(hand, straigtStart, straigtStart+5)
	} else {
		return false, 0
	}
}

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
func hasFlush(d Deck) (bool, Deck) {
	d.sort(-1)
	m := separateByFace(d)

	for _, v := range m {
		if v >= 5 {
			return true, d
		}
	}

	return false, d
}
