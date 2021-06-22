package main

import "fmt"

func main() {

}

func deckDemo() {

	cards := newDeck()
	cards.print()
}

func sliceDemo() {

	cards := deck{"One of Black", newCard()}
	deck.print(cards)

	cards = append(cards, "2 of Hearts")
	deck.print(cards)
}

func variableDemo() {

	//var card string = "Ace of Spades"

	card := "Ace of Spades"
	fmt.Println(card)

	card = newCard()
	fmt.Println(card)
}

func newCard() string {

	return "Five of Spades"
}
