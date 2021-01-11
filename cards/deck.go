package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

var deckSize int = 52

func newDeck() (d deck) {
	suits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	values := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range suits {
		for _, value := range values {
			d = append(d, value+" of "+suit)
		}
	}
	return d
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := 0; i < len(d); i++ {
		t := d[i]
		randomPosition := r.Intn(len(d))
		d[i] = d[randomPosition]
		d[randomPosition] = t
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func loadFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) deal(count int) (deck, deck) {
	hand := d[:count]
	remainingCards := d[count:]
	return hand, remainingCards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}
