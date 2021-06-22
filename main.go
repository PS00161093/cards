package main

import "fmt"

func main() {

	deckDemo()
}

func deckDemo() {

	cards := newDeck()
	hand, remainingHand := deal(cards, 5)
	hand.print()
	remainingHand.print()
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
