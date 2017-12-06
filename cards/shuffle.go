package cards

import (
	"math/rand"
	"time"
)

type ShuffleMethod interface {
	Shuffle(CardStack) []Card
}

type ReverseShuffle struct{}

func (r ReverseShuffle) Shuffle(cardStack CardStack) []Card {
	deckSize := len(cardStack.deck.cards)
	cards := make([]Card, deckSize)
	cardsIndex := 0
	for deckIndex := deckSize - 1; deckIndex > -1; deckIndex-- {
		cards[cardsIndex] = cardStack.deck.cards[deckIndex]
		cardsIndex++
	}

	return cards
}

type PerfectShuffle struct{}

func (p PerfectShuffle) Shuffle(cardStack CardStack) []Card {
	cards := []Card{}

	nextCardCount := GetNextCardCount()
	evenOdd := 2

	// Deal out a stack of cards randomly taking between 1 and 3 cards from top or bottom sequentially, do this 5 times
	for count := 0; count < 5; count++ {
		for cardStack.CardsLeft() > 0 {
			if evenOdd%2 == 0 {
				cards = append(cards, cardStack.DealCards(nextCardCount)...)
			} else {
				cards = append(cards, cardStack.DealCardsBottom(nextCardCount)...)
			}

			evenOdd++
			nextCardCount = GetNextCardCount()
		}

		cardStack.remainingCards = cards
		cards = []Card{}

	}

	return cardStack.remainingCards
}

func GetNextCardCount() int {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	return random.Intn(5) + 1
}
