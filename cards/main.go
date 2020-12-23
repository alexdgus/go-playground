package main

import "fmt"

//define a custom type
type animal struct {
	legs int
	tail bool
	name string
}

func main() {
	//define new variable
	var card1 string = "King of Clubs"
	//define new variable with type inferred
	var card2 = "King of Clubs"
	//define new variable alternate syntax
	card3 := "King of Clubs"
	//define new variable from function return value
	card4 := newCard()

	//define an array

	//define a slice with values
	cards1 := []string{newCard(), "Kind of Hearts", "King of Spades"}
	//define an empty slice
	cards2 := []string{}

	//append to a slice
	cards1 = append(cards1, "King of Diamonds")

	//iterate over cards
	for i, card := range cards1 {
		fmt.Println(card, i)
	}

	dog := new(animal)
	dog.legs = 4
	dog.tail = true
	dog.name = "Dog"

	d := deck{"King of Spades", "King of Clubs"}

	fmt.Println(card1)
	fmt.Println(card2)
	fmt.Println(card3)
	fmt.Println(card4)
	fmt.Println(cards1)
	fmt.Println(cards2)
	fmt.Println(dog)
	d.print()
}

func newCard() string {
	return "King of Clubs"
}
