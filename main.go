package main

import (
	"cards/cards"
	"fmt"
)

func main() {
	deck := cards.GetStandardDeck()
	cardStack := cards.NewCardStack(cards.PerfectShuffle{100}, deck, true)
	deviations := []float64{}
	for index := 0; index < 1000; index++ {
		cardStack.Shuffle()
		deviations = append(deviations, cardStack.GetDeviation())
	}

	fmt.Printf("Deviations:%+v\n", deviations)

	total := 0.0
	for _, dev := range deviations {
		total += dev
	}
	avg := total / float64(len(deviations))
	fmt.Printf("Average:%+v\n", avg)
}
