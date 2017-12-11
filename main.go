package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/bslezak/cards/cards"
)

func main() {
	// Get entropy from command line arg
	entropy, _ := strconv.Atoi(os.Args[1])

	// Run stats
	runStats(entropy)
}

// runStats creates a card deck and shuffles the deck 1000 times, collecting information on the statistical deviation of the deck
func runStats(entropy int) {

	// Create a standard deck of cards
	deck := cards.CreateStandardDeck()

	// Start shuffleTimes at 1, but we will shift this value incrementing by power of 2
	shuffleTimes := 1

	// Loop 8 times, incremeting shuffleTimes by power of two
	for interate := 0; interate < 8; interate++ {
		// Create a slice to collect stack deviations
		deviations := []float64{}

		// Get a cardStack that will be shuffled by NaturalShuffle
		cardStack := cards.NewCardStack(cards.NaturalShuffle{Shuffler: cards.Shuffler{ShuffleTimes: shuffleTimes, MaxEntropy: entropy}}, deck, false)

		for index := 0; index < 1000; index++ {
			cardStack.ResetStack()                                    // Reset the stack to reorder cards into their natural state
			cardStack.Shuffle()                                       // Shuffle the cards
			deviations = append(deviations, cardStack.GetDeviation()) // Calculate the deviation and collect that into our deviations slice
		}

		// Calculate the average deviation per 1000 shuffles
		avg := AvgDeviation(deviations)

		fmt.Printf("----\nShuffle Times:%+v\nEntropy:%+v\nAverage Dev:%+v\n\n", shuffleTimes, entropy, avg)

		// Shift shuffle times which increments by power of 2
		shuffleTimes = shuffleTimes << 1
	}
}

// AvgDeviation calculates and returns the average deviation for a population of float64 values
func AvgDeviation(deviations []float64) float64 {
	total := 0.0
	for _, dev := range deviations {
		total += dev
	}

	return total / float64(len(deviations))

}
