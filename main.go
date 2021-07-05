package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {

	cards := newDeck()
	cards.saveToFile("deck-of-cards")
}

func (d deck) toString() string {

	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {

	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
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
