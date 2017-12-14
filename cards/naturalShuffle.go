package cards

import (
	"math/rand"
	"time"
)

// NaturalShuffle is a Shuffler that attempts to shuffle cards in the most natural way possible
type NaturalShuffle struct {
	Shuffler
}

// Shuffle a card stack in the most natural way possible
// TODO: Improve this by not splitting the deck perfectly each time
func (shuffler NaturalShuffle) Shuffle(cardStack CardStack, shuffleTimes int) []Card {
	remainingCards := cardStack.remainingCards
	midpointOffset := (len(remainingCards) / 2) + GetMidPointOffset()

	for shuffleCount := 0; shuffleCount < shuffleTimes; shuffleCount++ {
		splitDeck := SplitDeck{Reverse(remainingCards[:midpointOffset]), remainingCards[midpointOffset:], 2}
		newCards := []Card{}
		nextCards := splitDeck.DealCards(shuffler.GetNextCardCount())
		for ; nextCards != nil; nextCards = splitDeck.DealCards(shuffler.GetNextCardCount()) {
			newCards = append(newCards, nextCards...)
		}
		remainingCards = newCards
		// fmt.Printf("Cards:%v\n\n", len(remainingCards))
	}

	return remainingCards
}

// GetMidPointOffset gets a signed integer between -3 and 4 to serve as an offset
func GetMidPointOffset() int {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	return random.Intn(7) - 3
}
