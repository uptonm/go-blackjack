package model

import "fmt"

type CardSuit string
type CardValue string

const (
	Hearts   CardSuit = "Hearts"
	Diamonds CardSuit = "Diamonds"
	Spades   CardSuit = "Spades"
	Clubs    CardSuit = "Clubs"
)

const (
	Ace   CardValue = "Ace"
	Two   CardValue = "2"
	Three CardValue = "3"
	Four  CardValue = "4"
	Five  CardValue = "5"
	Six   CardValue = "6"
	Seven CardValue = "7"
	Eight CardValue = "8"
	Nine  CardValue = "9"
	Ten   CardValue = "10"
	Jack  CardValue = "Jack"
	Queen CardValue = "Queen"
	King  CardValue = "King"
)

var CardSuits = [...]CardSuit{Hearts, Diamonds, Spades, Clubs}
var CardValues = [...]CardValue{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}

type Card struct {
	Suit  CardSuit  `json:"suit"`
	Value CardValue `json:"value"`
}

// String accepts a receiver of *Card and returns its printable name
func (card *Card) String() string {
	return fmt.Sprintf("%s of %s", card.Value, card.Suit)
}
