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
	want := Card{Suit: Spade, Rank: Ace}
	if cards[0] != want {
		t.Errorf("Something went wrong with DefaultSort! First card is: %v; Expected: %v", cards[0], want)
	}
}

func TestSort(t *testing.T) {
	cards := New(Sort(Less))
	want := Card{Suit: Spade, Rank: Ace}
	if cards[0] != want {
		t.Errorf("Something went wrong with Sort! First card is: %v; Expected: %v", cards[0], want)
	}
}

func TestJokers(t *testing.T) {
	want := 3
	cards := New(Jokers(want))
	count := 0
	for _, c := range cards {
		if c.Suit == Joker {
			count++
		}
	}
	if count != want {
		t.Errorf("Something went wrong! Not enough Jokers, wanted %d, got %d.", want, count)
	}
}

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}
	cards := New(Filter(filter))
	for _, c := range cards {
		if c.Rank == Two || c.Rank == Three {
			t.Errorf("Something went wrong! Expected all twos and threes to be filtered out.")
		}
	}
}
