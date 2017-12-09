package cards

import (
	"math/rand"
	"time"
)

// ShuffleMethod provides an interface to shuffling a stack of Cards
type ShuffleMethod interface {

	// Shuffle a stack of Cards
	Shuffle(CardStack) []Card
}

// PerfectShuffle
type PerfectShuffle struct {
	ShuffleTimes int
	MaxEntropy   int
}

// Shuffle attempts to shuffle cards by uniformly taking cards from top or bottom of the card stack. Entropy is introduced by a random number of cards taken each time
func (p PerfectShuffle) Shuffle(cardStack CardStack) []Card {
	cards := []Card{}

	nextCardCount := p.GetNextCardCount()
	evenOdd := 2

	// log.Println("Shuffling " + strconv.Itoa(p.ShuffleTimes) + " Times")
	// Deal out a stack of cards taking cards from top or bottom sequentially
	for count := 0; count < p.ShuffleTimes; count++ {
		for cardStack.CardsLeft() > 0 {
			if evenOdd%2 == 0 {
				cards = append(cards, cardStack.DealCards(nextCardCount)...)
			} else {
				cards = append(cards, cardStack.DealCardsBottom(nextCardCount)...)
			}

			evenOdd++
			nextCardCount = p.GetNextCardCount()
		}

		cardStack.remainingCards = cards
		cards = []Card{}

	}
	// fmt.Printf("Cardstack:%+v\n", cardStack.remainingCards)
	return cardStack.remainingCards
}

// Get the random number of cards that will be shuffled between 0 and n+1 times
func (p PerfectShuffle) GetNextCardCount() int {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	return random.Intn(p.MaxEntropy) + 1
}

// Get the random number of cards that will be shuffled between 0 and n+1 times
func (n NaturalShuffle) GetNextCardCount() int {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	return random.Intn(n.MaxEntropy) + 1
}

// NaturalShuffle
type NaturalShuffle struct {
	ShuffleTimes int
	MaxEntropy   int
}

// A SplitDeck is a stack of cards that is split in half
type SplitDeck struct {
	right              []Card
	left               []Card
	currentSideCounter int
}

// Deal cards from a split deck, taking from the right or left of the deck seqentially
func (splitDeck *SplitDeck) DealCards(count int) []Card {
	if splitDeck.currentSideCounter < 2 {
		splitDeck.currentSideCounter = 2
	}

	cards := []Card{}
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
	cards := []Card{}
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
	cards := []Card{}
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

// Shuffle a card stack in the most natual way possible
// TODO: Improve this by not splitting the deck perfectly each time
func (shuffler NaturalShuffle) Shuffle(cardStack CardStack) []Card {
	remainingCards := cardStack.remainingCards
	half := len(remainingCards) / 2

	for shuffleCount := 0; shuffleCount < shuffler.ShuffleTimes; shuffleCount++ {
		splitDeck := SplitDeck{Reverse(remainingCards[:half]), remainingCards[half:], 2}
		newCards := []Card{}
		nextCards := splitDeck.DealCards(shuffler.GetNextCardCount())
		for ; nextCards != nil; nextCards = splitDeck.DealCards(shuffler.GetNextCardCount()) {
			newCards = append(newCards, nextCards...)
		}
		remainingCards = newCards
		// fmt.Printf("Cards:%v\n\n", remainingCards)
	}

	return remainingCards
}

// Reverse a slice of Cards
func Reverse(cards []Card) []Card {
	count := len(cards)

	for index := 0; index < count/2; index++ {
		last := count - index - 1
		cards[index], cards[last] = cards[last], cards[index]
	}

	return cards
}
