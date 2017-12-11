package cards

// Card represents a playing card in the system
// Cards always have a suit and a rank, both integer values
type Card struct {
	suit int
	rank int
}

// SetSuit assigns the suit value of a card
func (card *Card) SetSuit(suit int) {
	card.suit = suit
}

// SetRank assigns the rank value of a card
func (card *Card) SetRank(rank int) {
	card.rank = rank
}

// NewCard initializes a new Card provided a suit and rank
func NewCard(suit int, rank int) Card {
	return Card{suit, rank}
}
