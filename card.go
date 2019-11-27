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

// New creates a new deck of cards and takes in functional options
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
	// sort.Slice sorts the given slice using the given function less(i, j int) bool
	// here we're calling another function, Less
	// could just do it like this:
	// sort.Slice(cards, func(i, j int) bool { return absRank(cards[i]) < absRank(cards[j]) }
	sort.Slice(cards, Less(cards))
	return cards
}

func Sort(less func(cards []Card) func(i, j int) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

// Less takes in a slice of cards and returns an anonymous less function
// need to create a closure so that the inner function can access the cards slice
func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}

func absRank(c Card) int {
	return int(c.Suit)*int(maxRank) + int(c.Rank)
}
