package main

import (
	"cards/cards"
	"fmt"
)

func main() {
	deck := cards.GetStandardDeck()
	fmt.Println(deck)
}
