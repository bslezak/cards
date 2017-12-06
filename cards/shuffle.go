package cards

import (
	"math/rand"
	"time"
)

type ShuffleMethod interface {
	Shuffle(Deck) []*Card
}

type ReverseShuffle struct{}

func (r ReverseShuffle) Shuffle(deck Deck) []*Card {
	deckSize := len(deck.cards)
	cards := make([]*Card, deckSize)
	cardsIndex := 0
	for deckIndex := deckSize - 1; deckIndex > -1; deckIndex-- {
		cards[cardsIndex] = &deck.cards[deckIndex]
		cardsIndex++
	}

	return cards
}

type PerfectShuffle struct{}

func (p PerfectShuffle) Shuffle(deck Deck) []*Card {
	deckSize := len(deck.cards)
	cards := make([]*Card, deckSize)
	nextCardCount := GetNextCardCount()

	// Iterate the deck by nextCardCount, randomizing card assignment
}

func GetNextCardCount() int {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	return random.Intn(3) + 1
}
