package cards

// PerfectShuffle is a shuffle that attempts to shuffle cards in a very uniform manner
type PerfectShuffle struct {
	Shuffler
}

// Shuffle attempts to shuffle cards by uniformly taking cards from top or bottom of the card stack.
// Entropy is introduced by a random number of cards taken each time
func (shuffler PerfectShuffle) Shuffle(cardStack CardStack, shuffleTimes int) []Card {
	cards := []Card{}

	nextCardCount := shuffler.GetNextCardCount()
	evenOdd := 2

	// log.Println("Shuffling " + strconv.Itoa(p.ShuffleTimes) + " Times")
	// Deal out a stack of cards taking cards from top or bottom sequentially
	for count := 0; count < shuffleTimes; count++ {
		for cardStack.CardsLeft() > 0 {
			if evenOdd%2 == 0 {
				cards = append(cards, cardStack.DealCards(nextCardCount)...)
			} else {
				cards = append(cards, cardStack.DealCardsBottom(nextCardCount)...)
			}

			evenOdd++
			nextCardCount = shuffler.GetNextCardCount()
		}

		cardStack.remainingCards = cards
		cards = []Card{}

	}
	// fmt.Printf("Cardstack:%+v\n", cardStack.remainingCards)
	return cardStack.remainingCards
}
