package cards

import (
	"log"
	"math/rand"
	"strconv"
	"time"
)

type ShuffleMethod interface {
	Shuffle(CardStack) []Card
}

type ReverseShuffle struct {
	ShuffleTimes int
}

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

type PerfectShuffle struct {
	ShuffleTimes int
}

func (p PerfectShuffle) Shuffle(cardStack CardStack) []Card {
	cards := []Card{}

	nextCardCount := GetNextCardCount()
	evenOdd := 2

	log.Println("Shuffling " + strconv.Itoa(p.ShuffleTimes) + " Times")
	// Deal out a stack of cards randomly taking between 1 and 3 cards from top or bottom sequentially, do this 5 times
	for count := 0; count < p.ShuffleTimes; count++ {
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
	// fmt.Printf("Cardstack:%+v\n", cardStack.remainingCards)
	return cardStack.remainingCards
}

func GetNextCardCount() int {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	return random.Intn(5) + 1
}
