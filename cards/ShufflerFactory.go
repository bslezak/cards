package cards

// BuildShuffler performs the building of shufflers
func BuildShuffler(shufflerName string, shuffleTimes int, maxEntropy int) ShuffleMethod {
	var shuffler ShuffleMethod
	switch shufflerName {
	case "perfect":
		shuffler = PerfectShuffle{Shuffler{ShuffleTimes: shuffleTimes, MaxEntropy: maxEntropy}}

	case "natural":
		shuffler = NaturalShuffle{Shuffler{ShuffleTimes: shuffleTimes, MaxEntropy: maxEntropy}}
	}

	return shuffler
}
