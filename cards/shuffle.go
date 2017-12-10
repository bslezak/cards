package cards

// ShuffleMethod provides an interface to shuffling a stack of Cards
type ShuffleMethod interface {

	// Shuffle a stack of Cards
	Shuffle(CardStack) []Card
}

// Shuffler retains static data for a ShuffleMethod
type Shuffler struct {
	ShuffleTimes int
	MaxEntropy   int
}

// A SplitDeck is a stack of cards that is split in half
type SplitDeck struct {
	right              []Card
	left               []Card
	currentSideCounter int
}

// DealCards deals cards from a split deck, taking from the right or left of the deck sequentially
func (splitDeck *SplitDeck) DealCards(count int) []Card {
	if splitDeck.currentSideCounter < 2 {
		splitDeck.currentSideCounter = 2
	}

	var cards []Card
	if splitDeck.currentSideCounter%2 == 0 {
		cards = splitDeck.dealCardsRight(count)
	} else {
		cards = splitDeck.dealCardsLeft(count)
	}
	splitDeck.currentSideCounter++

	return cards

}

// Deal cards from the right side of the deck
func (splitDeck *SplitDeck) dealCardsRight(count int) []Card {
	var cards []Card
	if splitDeck.right != nil {
		if len(splitDeck.right) > count {
			cards = splitDeck.right[:count]
			splitDeck.right = splitDeck.right[count:]
		} else {
			cards = splitDeck.right
			cards = append(cards, splitDeck.left...)
			splitDeck.right = nil
			splitDeck.left = nil
		}
	} else {
		cards = nil
	}

	return cards
}

// Deal cards from the left side of the deck
func (splitDeck *SplitDeck) dealCardsLeft(count int) []Card {
	var cards []Card
	if splitDeck.left != nil {
		if len(splitDeck.left) > count {
			cards = splitDeck.left[:count]
			splitDeck.left = splitDeck.left[count:]
		} else {
			cards = splitDeck.left
			cards = append(cards, splitDeck.right...)
			splitDeck.left = nil
			splitDeck.right = nil
		}
	} else {
		cards = nil
	}

	return cards
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
