package cards

// A Deck of Cards
type Deck struct {
	cards []Card
}

// Get the Cards
func (deck Deck) GetCards() []Card {
	return deck.cards
}

// Create a standard deck of cards, 4 suits and 13 ranks
func CreateStandardDeck() Deck {
	deck := Deck{}

	for suitNum := 1; suitNum < 5; suitNum++ {
		for index := 1; index < 14; index++ {
			deck.cards = append(deck.cards, Card{suitNum, index})
		}
	}

	return deck
}

// Convert suit integer to a string
func (c Card) StandardSuit() string {
	suit := []string{"", "Clubs", "Hearts", "Diamonds", "Spades"}

	return suit[c.suit]
}

// Convert rank integer to a string
func (c Card) StandardRank() string {
	rank := []string{"", "Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	return rank[c.rank]
}

// Get the textual representation of a card
func (card Card) GetFace() string {
	return card.StandardRank() + " of " + card.StandardSuit()
}
