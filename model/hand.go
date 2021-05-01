package model

import (
	"fmt"
	"strconv"
)

type Hand struct {
	Dealer bool   `json:"dealer"`
	Cards  []Card `json:"cards"`
	Value  int    `json:"value"`
}

// InitHand accepts a boolean stating whether the user is the dealer, and returns an empty Hand
func InitHand(dealer bool) *Hand {
	return &Hand{
		Dealer: dealer,
		Cards:  []Card{},
		Value:  0,
	}
}

// AddCard accepts a receiver of *Hand and a parameter of Card and appends the Card to the player's hand
func (hand *Hand) AddCard(card Card) {
	hand.Cards = append(hand.Cards, card)
	hand.CalculateValue()
}

// CalculateValue accepts a receiver of *Hand, calculates, and sets the value of its cards
func (hand *Hand) CalculateValue() {
	value := 0
	hasAce := false
	for _, card := range hand.Cards {
		numericValue, err := strconv.Atoi(string(card.Value))
		if err == nil {
			value += numericValue
		} else {
			if card.Value == Ace {
				value += 11
				hasAce = true
			} else {
				value += 10
			}
		}
	}
	if hasAce && value > 21 {
		value -= 10
	}

	hand.Value = value
}

// GetValue accepts a receiver of *Hand and calculates its value in place, returning the value
func (hand *Hand) GetValue() int {
	hand.CalculateValue()
	return hand.Value
}

// Display accepts a receiver of *Hand and displays its value, hiding the first card if Hand.Dealer = true
// and gameOver = false
func (hand *Hand) Display(gameOver bool) {
	if hand.Dealer && !gameOver {
		fmt.Printf("*hidden*\n%s\n", hand.Cards[1].String())
		return
	}

	for _, card := range hand.Cards {
		fmt.Printf("%s\n", card.String())
	}
	fmt.Printf("Value: %d\n", hand.GetValue())
}

// HasBlackjack accepts a receiver of *Hand and returns true if the *Hand.GetValue() is 21
func (hand *Hand) HasBlackjack() bool {
	return hand.GetValue() == 21
}

// HasBust is a helper method accepting a receiver of *Hand and returning true if GetValue > 21
func (hand *Hand) HasBust() bool {
	return hand.GetValue() > 21
}