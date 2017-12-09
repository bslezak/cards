package cards

import (
	"math"
)

// A stack of cards that can be shuffled and dealt
type CardStack struct {
	deck           Deck
	shuffler       ShuffleMethod
	remainingCards []Card
}

// Create a new card stack from a deck and shuffle method.
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

// Shuffle the card stack
func (cardStack *CardStack) Shuffle() {
	cardStack.ResetStack()
	cardStack.remainingCards = cardStack.shuffler.Shuffle(*cardStack)
}

// Reset a card stack, setting it back to it's original state prior to shuffling
func (cardStack *CardStack) ResetStack() {
	deckSize := len(cardStack.deck.cards)
	cardStack.remainingCards = make([]Card, deckSize)
	copy(cardStack.remainingCards, cardStack.deck.cards)
}

// Deal cards from a cardstack
func (cardStack *CardStack) DealCards(count int) []Card {
	cards := []Card{}
	if count < cardStack.CardsLeft() {
		cards = cardStack.remainingCards[:count]
		cardStack.remainingCards = cardStack.remainingCards[count:]
	} else {
		cards = cardStack.remainingCards
		cardStack.remainingCards = []Card{}
	}
	// log.Println("Cards Remaining:" + strconv.Itoa(len(cardStack.remainingCards)))
	return cards
}

// Deal cards from the bottom of a cardstack
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
	// log.Println("Cards Remaining:" + strconv.Itoa(len(cardStack.remainingCards)))
	return cards
}

// Get the number of cards left in a stack
func (cardStack CardStack) CardsLeft() int {
	return len(cardStack.remainingCards)
}

// Get the deviation of a cardstack in it's current state.
// Deviation is determined by one cards proximity to another from it's unshuffled state. An unshuffled card stack would have a deviation of 0
// The maximum deviation a card may have is 22, meaning the card next to it is of a different suit (+10) and is the lowest ranking card
// next to the highest ranking card
func (cardStack CardStack) GetDeviation() float64 {
	deviations := []float64{}

	for count := 0; count < 51; count++ {
		card1 := cardStack.remainingCards[count]
		card2 := cardStack.remainingCards[count+1]
		diff := 0.0
		if card1.suit != card2.suit {
			diff += 10
		}
		diff += math.Abs(float64(card1.rank-card2.rank)) + 13.0
		deviations = append(deviations, diff)
	}
	// fmt.Printf("Deviations:%+v\n", deviations)
	product := float64(0)
	for _, dev := range deviations {
		product += math.Abs(dev) - 1
	}
	return product
}
