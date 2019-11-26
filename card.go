package deck

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

