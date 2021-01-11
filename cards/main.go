package main

func main() {
	d := newDeck()
	d.shuffle()
	d.print()
	//h, remainingDeck := d.deal(5)

	//remainingDeck.shuffle()
	//remainingDeck.print()
	//h.print()
	//fmt.Println(d.toString())
	//remainingDeck.saveToFile("drawPile")
	//newd := loadFromFile("drawPile")
	//newd.print()
}
