//go:generate stringer -type=Suit,Rank

package deck

import (
	"fmt"
	"sort"
)

type Card struct {
	Suit
	Rank
}

// make it a unit8 because there aren't a ton of suits
// not a big deal tho
type Suit uint8

const (
	// iota starts at 0, increments for each following variable
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

// okay to make it an array because it won't change
var suits = [...]Suit{Spade, Diamond, Club, Heart}

type Rank uint8

const (
	// skip zero so numbers match up
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

// makes it simpler when creating a new deck
const (
	minRank = Ace
	maxRank = King
)

// add a Stringer so suits don't appear as numbers
func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

// New creates a new deck of cards
func New(opts ...func([]Card) []Card) []Card {
	var cards []Card

	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}
	for _, opt := range opts {
		cards = opt(cards)
	}

	return cards
}

func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}

func absRank(c Card) int {
	return int(c.Suit)*int(maxRank) + int(c.Rank)
}
