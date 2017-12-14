package cards

// BuildShuffler performs the building of shufflers
func BuildShuffler(shufflerName string, maxEntropy int) ShuffleMethod {
	var shuffler ShuffleMethod
	switch shufflerName {
	case "perfect":
		shuffler = PerfectShuffle{Shuffler{MaxEntropy: maxEntropy}}

	case "natural":
		shuffler = NaturalShuffle{Shuffler{MaxEntropy: maxEntropy}}
	}

	return shuffler
}
