package cards

// Deck is a deck of Cards
type Deck struct {
	cards []Card
}

// GetCards returns all the Cards from the deck
func (deck Deck) GetCards() []Card {
	return deck.cards
}

// CreateStandardDeck creates a standard deck of cards, 4 suits and 13 ranks
func CreateStandardDeck() Deck {
	deck := Deck{}

	for suitNum := 1; suitNum < 5; suitNum++ {
		for index := 1; index < 14; index++ {
			deck.cards = append(deck.cards, Card{suitNum, index})
		}
	}

	return deck
}

// StandardSuit converts suit integer to a string
func (card Card) StandardSuit() string {
	suit := []string{"", "Clubs", "Hearts", "Diamonds", "Spades"}

	return suit[card.suit]
}

// StandardRank converts rank integer to a string
func (card Card) StandardRank() string {
	rank := []string{"", "Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	return rank[card.rank]
}

// GetFace gets the textual representation of a card
func (card Card) GetFace() string {
	return card.StandardRank() + " of " + card.StandardSuit()
}
