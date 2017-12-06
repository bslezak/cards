package main

import (
	"cards/cards"
	"fmt"
)

func main() {
	deck := cards.GetStandardDeck()
	cardStack := cards.NewCardStack(cards.PerfectShuffle{}, deck, true)
	fmt.Printf("%+v\n", cardStack)

}
