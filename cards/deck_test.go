package main

import (
	"os"
	"testing"
)

var testFile string = "_deck_testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != deckSize {
		t.Errorf("Expected length of %v but got %v", deckSize, len(d))
	}

	if d[0] != "Ace of Hearts" {
		t.Errorf("Expected Ace of Hearts but got %v", d[0])
	}

	if d[len(d)-1] != "King of Spades" {
		t.Errorf("Expected King of Spades but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove(testFile)
	d := newDeck()
	d.saveToFile(testFile)
	loadedDeck := loadFromFile(testFile)
	if len(loadedDeck) != deckSize {
		t.Errorf("Expected %v cards but got %v", deckSize, len(loadedDeck))
	}
	os.Remove(testFile)
}
