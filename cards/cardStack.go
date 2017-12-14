package cards

import (
	"math"
)

// CardStack stack of cards that can be shuffled and dealt
type CardStack struct {
	deck           Deck
	shuffler       ShuffleMethod
	remainingCards []Card
}

// NewCardStack instantiates a new card stack provided a deck and shuffle method.
// The parameter instantShuffle is a flag indicating whether the CardStack should be shuffled immediately after creation
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

// Shuffle shuffles the cards in a cardstack by calling ShufflerMethod.Shuffle
func (cardStack *CardStack) Shuffle() {
	cardStack.ResetStack()
	cardStack.remainingCards = cardStack.shuffler.Shuffle(*cardStack)
}

// ResetStack returns a card stack back to it's original deck state prior to any shuffling
func (cardStack *CardStack) ResetStack() {
	deckSize := len(cardStack.deck.cards)
	cardStack.remainingCards = make([]Card, deckSize)
	copy(cardStack.remainingCards, cardStack.deck.cards)
}

// DealCards deals a specified number of cards from the top of a cardstack
func (cardStack *CardStack) DealCards(cardCount int) []Card {
	var cards []Card
	if cardCount < cardStack.CardsLeft() {
		cards = cardStack.remainingCards[:cardCount]
		cardStack.remainingCards = cardStack.remainingCards[cardCount:]
	} else {
		cards = cardStack.remainingCards
		cardStack.remainingCards = nil
	}
	// log.Println("Cards Remaining:" + strconv.Itoa(len(cardStack.remainingCards)))
	return cards
}

// DealCardsBottom deals cards from the bottom of a cardstack
func (cardStack *CardStack) DealCardsBottom(cardCount int) []Card {
	var cards []Card
	if cardCount < cardStack.CardsLeft() {
		lastCard := len(cardStack.remainingCards)
		cards = cardStack.remainingCards[lastCard-cardCount:]
		cardStack.remainingCards = cardStack.remainingCards[:lastCard-cardCount]
	} else {
		cards = cardStack.remainingCards
		cardStack.remainingCards = []Card{}
	}
	// log.Println("Cards Remaining:" + strconv.Itoa(len(cardStack.remainingCards)))
	return cards
}

// CardsLeft returns the number of cards left in a stack
func (cardStack CardStack) CardsLeft() int {
	return len(cardStack.remainingCards)
}

// GetDeviation get the deviation of a cardstack in it's current state.
// Deviation is determined by one cards proximity to another from it's unshuffled state. An unshuffled card stack would have a deviation of 0
// The maximum deviation a card may have is 22, meaning the card next to it is of a different suit (+10) and is the lowest ranking card
// next to the highest ranking card
func (cardStack CardStack) GetDeviation() float64 {
	deviations := []float64{}

	for cardCount := 0; cardCount < 51; cardCount++ {
		card1 := cardStack.remainingCards[cardCount]
		card2 := cardStack.remainingCards[cardCount+1]
		diff := 0.0

		if card2.suit-1 != card1.suit && card2.rank-card1.rank != 13 {
			if card1.suit != card2.suit {
				diff += math.Abs(float64(card2.suit-card1.suit)) * 10
			}

			diff += math.Abs(float64(card2.rank-card1.rank) - 1)
		}

		deviations = append(deviations, diff)
	}
	// fmt.Printf("Deviations:%+v\n", deviations)
	product := float64(0)
	for _, dev := range deviations {
		product += math.Abs(dev)
	}
	return product
}
