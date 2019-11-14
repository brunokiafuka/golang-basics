package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	// t is our test handler
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Ace of Spades, but got %v", d[len(d)-1])
	}
}

func TestSaveReadToFileDeck(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDecck := newDeckFromFile("_decktesting")

	if len(loadedDecck) != 16 {
		t.Errorf("Expected loaded deck length of 16, but got %v", len(d))
	}

	os.Remove("_decktesting")
}
