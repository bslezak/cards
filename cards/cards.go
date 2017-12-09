package cards

// A Card
type Card struct {
	suit int
	rank int
}

// Set the suit of a Card
func (card *Card) SetSuit(suit int) {
	card.suit = suit
}

// Set the rank of a card
func (card *Card) SetRank(rank int) {
	card.rank = rank
}

// Create a new card provided a suit and rank
func NewCard(suit int, rank int) Card {
	return Card{suit, rank}
}
