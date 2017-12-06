package cards

import (
	"math"
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

func (cardStack *CardStack) Shuffle() {
	cardStack.ResetStack()
	cardStack.remainingCards = cardStack.shuffler.Shuffle(*cardStack)
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
	// log.Println("Cards Remaining:" + strconv.Itoa(len(cardStack.remainingCards)))
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
	// log.Println("Cards Remaining:" + strconv.Itoa(len(cardStack.remainingCards)))
	return cards
}

func (cardStack CardStack) CardsLeft() int {
	return len(cardStack.remainingCards)
}

func (cardStack CardStack) GetDeviation() float64 {
	deviations := []float64{}

	for count := 0; count < 51; count++ {
		card1 := cardStack.remainingCards[count]
		card2 := cardStack.remainingCards[count+1]
		diff := float64(card1.suit-card2.suit) * 10
		diff = diff + float64(card1.rank-card2.rank)
		deviations = append(deviations, diff)
	}
	// fmt.Printf("Deviations:%+v\n", deviations)
	product := float64(0)
	for _, dev := range deviations {
		product += math.Abs(dev) - 1
	}
	return product
}
