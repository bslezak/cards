package cards

import (
	"math/rand"
	"time"
)

// ShuffleMethod provides an interface to shuffling a stack of Cards
type ShuffleMethod interface {

	// Shuffle a stack of Cards, returning a slice of Cards
	Shuffle(CardStack) []Card
}

// Shuffler retains static data for a ShuffleMethod
type Shuffler struct {
	ShuffleTimes int
	MaxEntropy   int
}

// Reverse reverses the order of a slice of Cards
func Reverse(cards []Card) []Card {
	count := len(cards)

	for index := 0; index < count/2; index++ {
		last := count - index - 1
		cards[index], cards[last] = cards[last], cards[index]
	}

	return cards
}

// GetNextCardCount returns a random integer between 1 and MaxEntropy+1
func (shuffler Shuffler) GetNextCardCount() int {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	return random.Intn(shuffler.MaxEntropy) + 1
}
