package main

import "testing"
import "os"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("expected deck length is 16 but got %v", len(d))
	}
}

func TestsaveToFileAndnewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("expected 16 cards in deck, but got %v", len(loadedDeck))

	}
	os.Remove("_decktesting")
}
