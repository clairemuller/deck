package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Ten, Suit: Spade})
	fmt.Println(Card{Suit: Joker})

	// Output:
	// Ace of Hearts
	// Ten of Spades
	// Joker
}

func TestNew(t *testing.T) {
	cards := New()
	if len(cards) != 52 {
		t.Errorf("Wrong number of cards in a new deck! Got %d, wanted 52.", len(cards))
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	want := Card{Suit: Spade, Rank: Three}
	if cards[0] != want {
		t.Errorf("Something went wrong with DefaultSort! First card is: %v; Expected: %v", cards[0], want)
	}
}
