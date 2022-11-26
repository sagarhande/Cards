package main

import "testing"

func TestNewDeck(t *testing.T) {

	d := newDeck()

	if len(d) != 40 {
		t.Errorf("Expected length of deck 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be `Ace of Spades`, but got `%v`", d[0])
	}
	if d[len(d)-1] != "Ten of Clubs" {
		t.Errorf("Expected last card to be `Ten of Clubs`, but got `%v`", d[len(d)-1])
	}
}

func TestDeal(t *testing.T) {
	d := newDeck()
	n := len(d)
	hand, remaining := deal(d, 4)

	if len(hand) != 4 {
		t.Errorf("Deal should return slice of length 4, but got %v", len(hand))
	}
	if len(remaining) != n-4 {
		t.Errorf("Deal should return slice of length %v, but got %v", n-4, len(remaining))
	}
}
