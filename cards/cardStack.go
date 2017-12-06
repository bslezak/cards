package cards

type CardStack struct {
	cardDeck       Deck
	shuffler       ShuffleMethod
	remainingCards []*Card
}

func NewCardStack(shuffler ShuffleMethod, deck Deck, instantShuffle bool) CardStack {
	stack := CardStack{}
	stack.cardDeck = deck
	stack.shuffler = shuffler

	if instantShuffle {
		stack.remainingCards = stack.shuffler.Shuffle(stack.cardDeck)
	}

	return stack
}

func (cardStack CardStack) Shuffle() {
	cardStack.remainingCards = cardStack.shuffler.Shuffle(cardStack.cardDeck)
}
