package cards

import (
	"log"
	"strconv"
)

type CardStack struct {
	deck           Deck
	shuffler       ShuffleMethod
	remainingCards []Card
}

func NewCardStack(shuffler ShuffleMethod, deck Deck, instantShuffle bool) CardStack {
	stack := CardStack{}
	stack.deck = deck
	stack.remainingCards = stack.deck.GetCards()
	stack.shuffler = shuffler

	if instantShuffle {
		stack.remainingCards = stack.shuffler.Shuffle(stack)
	}

	return stack
}

func (cardStack CardStack) Shuffle() {
	cardStack.ResetStack()
	cardStack.remainingCards = cardStack.shuffler.Shuffle(cardStack)
}

func (cardStack *CardStack) ResetStack() {
	deckSize := len(cardStack.deck.cards)
	cardStack.remainingCards = make([]Card, deckSize)
	copy(cardStack.remainingCards, cardStack.deck.cards)
}

func (cardStack *CardStack) DealCards(count int) []Card {
	cards := []Card{}
	if count < cardStack.CardsLeft() {
		cards = cardStack.remainingCards[:count]
		cardStack.remainingCards = cardStack.remainingCards[count:]
	} else {
		cards = cardStack.remainingCards
		cardStack.remainingCards = []Card{}
	}
	log.Println("Cards Remaining:" + strconv.Itoa(len(cardStack.remainingCards)))
	return cards
}

func (cardStack *CardStack) DealCardsBottom(count int) []Card {
	cards := []Card{}
	if count < cardStack.CardsLeft() {
		lastCard := len(cardStack.remainingCards)
		cards = cardStack.remainingCards[lastCard-count:]
		cardStack.remainingCards = cardStack.remainingCards[:lastCard-count]
	} else {
		cards = cardStack.remainingCards
		cardStack.remainingCards = []Card{}
	}
	log.Println("Cards Remaining:" + strconv.Itoa(len(cardStack.remainingCards)))
	return cards
}

func (cardStack CardStack) CardsLeft() int {
	return len(cardStack.remainingCards)
}
