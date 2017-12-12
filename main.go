package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/bslezak/cards/cards"
	"github.com/fatih/color"
)

func main() {

	// Check command line args
	if len(os.Args[1:]) > 0 {
		// Get entropy from command line arg
		entropy, errors := strconv.Atoi(os.Args[1])

		if errors == nil {
			naturalFlag := flag.Bool("natural", false, "Use natural shuffling method")
			perfectFlag := flag.Bool("perfect", false, "Use perfect shuffling method")
			flag.Parse()
			if !*naturalFlag && !*perfectFlag {
				PrintError("One of --natural or --perfect options should be set")
			}
			// Run stats
			runStats(entropy)
		} else {
			PrintError("No entropy provided. Please provide an integer to seed entropy")
		}
	} else {
		PrintHelp()
	}

}

// PrintError prints an error, the help, and exits
func PrintError(error string) {
	color.Red(fmt.Sprintf("\nError: %s\n\n", error))
	PrintHelp()
	os.Exit(1)
}

// PrintHelp prints help and usage
func PrintHelp() {
	desc := `
Cards runs simulations of card shuffling and prints statistics

Usage:
	cards.exe <entropy> (--natural | --perfect)
	
entropy	An unsigned integer. This is used during shuffling to determine how many random cards are chosen between 0 and value

Options:
	--natural	Use natural shuffling method
	--perfect	Use perfect shuffling method
	`
	fmt.Println(desc)
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
