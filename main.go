package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/bslezak/cards/cards"
	"github.com/fatih/color"
	"gonum.org/v1/gonum/stat"
)

func main() {

	// Check command line inputs
	entropy, shufflerName := CheckInputs(os.Args)

	// All clear so continue

	// Run stats
	runStats(entropy, shufflerName)
}

// CheckInputs quality checks command line inputs
func CheckInputs(args []string) (int, string) {

	naturalFlag := flag.Bool("natural", false, "Use natural shuffling method")
	perfectFlag := flag.Bool("perfect", false, "Use perfect shuffling method")
	help := flag.Bool("help", false, "Display help and usage")
	flag.Parse()

	// If help, print it and exit
	if *help {
		PrintHelp()
		os.Exit(0)
	}
	// Get entropy from command line arg
	entropy, errors := strconv.Atoi(os.Args[2])

	// If error converting argument to integer, print errors and exit
	if errors != nil {
		PrintError("No entropy provided. Please provide an integer to seed entropy")
	}

	// If neither flag is set, print error and exit
	if !*naturalFlag && !*perfectFlag {
		PrintError("One of --natural or --perfect options should be set")
	}

	var shufflerName string
	if *naturalFlag {
		shufflerName = "natural"
	} else {
		shufflerName = "perfect"
	}

	return entropy, shufflerName
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
	cards.exe (--natural | --perfect) <entropy> 
	
<entropy>	An unsigned integer. This is used during shuffling to determine how many random cards are chosen between 0 and value

Options:
	--natural	Use natural shuffling method
	--perfect	Use perfect shuffling method
	`
	fmt.Println(desc)
}

// runStats creates a card deck and shuffles the deck 1000 times, collecting information on the statistical deviation of the deck
func runStats(entropy int, shufflerName string) {

	// Start shuffleTimes at 1, but we will shift this value incrementing by power of 2
	shuffleTimes := 1

	// Create a cardStack that will be shuffled
	deck := cards.CreateStandardDeck()
	shuffler := cards.BuildShuffler(shufflerName, entropy)
	cardStack := cards.NewCardStack(deck, shuffler, shuffleTimes, false)

	// Loop 8 times, incremeting shuffleTimes by power of two
	for iterate := 0; iterate < 8; iterate++ {
		// Create a slice to collect stack deviations
		deviations := []float64{}

		cardStack.ShuffleTimes = shuffleTimes
		for index := 0; index < 1000; index++ {
			cardStack.ResetStack()                                    // Reset the stack to reorder cards into their natural state
			cardStack.Shuffle()                                       // Shuffle the cards
			deviations = append(deviations, cardStack.GetDeviation()) // Calculate the deviation and collect that into our deviations slice
		}

		// Calculate the average deviation per 1000 shuffles
		avg := AvgDeviation(deviations)
		variance := stat.Variance(deviations, nil)
		fmt.Printf("----\nShuffle Times:%+v\nEntropy:%+v\nAverage Dev:%+v\nStandard Deviation:%+v\n\n", shuffleTimes, entropy, avg, variance)

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
