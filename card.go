//go:generate stringer -type=Suit,Rank

package deck

import "fmt"

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

// add a Stringer so suits don't appear as numbers
func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}
