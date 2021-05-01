package model

import (
	"errors"
	"math/rand"
	"time"
)

type Deck struct {
	Cards []Card `json:"cards"`
}

// Shuffle accepts a receiver of a *Deck and shuffles it in place
func (deck *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck.Cards), func(i, j int) {
		deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i]
	})
}

// Deal accepts a receiver of a *Deck and removes and returns the card at index 0
func (deck *Deck) Deal() (*Card, error) {
	if len(deck.Cards) == 0 {
		return nil, errors.New("deck_empty")
	}

	var card *Card
	card, deck.Cards = &deck.Cards[0], deck.Cards[1:]

	return card, nil
}

// InitDeck utilizes the constants of CardSuits and CardValues to return a shuffled *Deck
func InitDeck() *Deck {
	cards := make([]Card, 0)

	for _, suit := range CardSuits {
		for _, value := range CardValues {
			cards = append(cards, Card{
				Suit:  suit,
				Value: value,
			})
		}
	}

	deck := Deck{
		Cards: cards,
	}

	deck.Shuffle()

	return &deck
}
