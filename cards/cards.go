package cards

type Card struct {
	suit int
	rank int
}

func (card *Card) SetSuit(suit int) {
	card.suit = suit
}

func (card *Card) SetRank(rank int) {
	card.rank = rank
}

func NewCard(suit int, rank int) Card {
	return Card{suit, rank}
}
