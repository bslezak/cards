package cards

type Deck struct {
	cards []Card
}

func (deck Deck) GetCards() []Card {
	return deck.cards
}

func GetStandardDeck() Deck {
	deck := Deck{}

	for suitNum := 1; suitNum < 5; suitNum++ {
		for index := 1; index < 14; index++ {
			deck.cards = append(deck.cards, Card{suitNum, index})
		}
	}

	return deck
}

func (c Card) StandardSuit() string {
	suit := []string{"", "Clubs", "Hearts", "Diamonds", "Spades"}

	return suit[c.suit]
}

func (c Card) StandardRank() string {
	rank := []string{"", "Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	return rank[c.rank]
}

func (card Card) GetFace() string {
	return card.StandardRank() + " of " + card.StandardSuit()
}