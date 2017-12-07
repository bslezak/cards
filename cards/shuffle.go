package cards

import (
	"math/rand"
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
	MaxEntropy   int
}

func (p PerfectShuffle) Shuffle(cardStack CardStack) []Card {
	cards := []Card{}

	nextCardCount := p.GetNextCardCount()
	evenOdd := 2

	// log.Println("Shuffling " + strconv.Itoa(p.ShuffleTimes) + " Times")
	// Deal out a stack of cards randomly taking between 1 and 3 cards from top or bottom sequentially, do this 5 times
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

func (p PerfectShuffle) GetNextCardCount() int {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	return random.Intn(p.MaxEntropy) + 1
}

func (n NaturalShuffle) GetNextCardCount() int {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	return random.Intn(n.MaxEntropy) + 1
}

type NaturalShuffle struct {
	ShuffleTimes int
	MaxEntropy   int
}

type SplitDeck struct {
	right              []Card
	left               []Card
	currentSideCounter int
}

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

func Reverse(cards []Card) []Card {
	count := len(cards)

	for index := 0; index < count/2; index++ {
		last := count - index - 1
		cards[index], cards[last] = cards[last], cards[index]
	}

	return cards
}
