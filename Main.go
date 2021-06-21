package main

import "fmt"

func main() {

	//var card string = "Ace of Spades"

	card := "Ace of Spades"
	fmt.Println(card)

	card = newCard()
	fmt.Println(card)
}

func newCard() string {

	return "Five of Spades"
}
