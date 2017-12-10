package cards

import (
	"math/rand"
	"time"
)

// GetNextCardCount returns a random integer between 1 and MaxEntropy+1
func (shuffler NaturalShuffle) GetNextCardCount() int {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	return random.Intn(shuffler.MaxEntropy) + 1
}

// NaturalShuffle is a Shuffler that attempts to shuffle cards in the most natural way possible
type NaturalShuffle struct {
	*Shuffler
}

// Shuffle a card stack in the most natural way possible
// TODO: Improve this by not splitting the deck perfectly each time
func (shuffler NaturalShuffle) Shuffle(cardStack CardStack) []Card {
	remainingCards := cardStack.remainingCards
	half := len(remainingCards) / 2

	for shuffleCount := 0; shuffleCount < shuffler.ShuffleTimes; shuffleCount++ {
		splitDeck := SplitDeck{Reverse(remainingCards[:half]), remainingCards[half:], 2}
		newCards := []Card{}
		nextCards := splitDeck.DealCards(shuffler.GetNextCardCount())
		for ; nextCards != nil; nextCards = splitDeck.DealCards(shuffler.GetNextCardCount()) {
			newCards = append(newCards, nextCards...)
		}
		remainingCards = newCards
		// fmt.Printf("Cards:%v\n\n", remainingCards)
	}

	return remainingCards
}
