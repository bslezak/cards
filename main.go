package main

import (
	"cards/cards"
	"fmt"
	"os"
	"strconv"
)

func main() {
	entropy, _ := strconv.Atoi(os.Args[1])
	runStats(entropy)
}

func runStats(entropy int) {

	deck := cards.GetStandardDeck()
	shuffleTimes := 1

	for interate := 0; interate < 8; interate++ {
		cardStack := cards.NewCardStack(cards.NaturalShuffle{shuffleTimes, entropy}, deck, true)
		deviations := []float64{}
		for index := 0; index < 1000; index++ {
			cardStack.Shuffle()
			deviations = append(deviations, cardStack.GetAvgDev())
		}

		total := 0.0
		for _, dev := range deviations {
			total += dev
		}
		avg := total / float64(len(deviations))
		fmt.Printf("----\nShuffle Times:%+v\nEntropy:%+v\nAverage Dev:%+v\n\n", shuffleTimes, entropy, avg)
		shuffleTimes = shuffleTimes << 1
	}
}
