package main

import "fmt"

func main() {

	sliceDemo()
}

func sliceDemo() {

	cards := []string{"One of Black", newCard()}
	printSlice(cards)

	cards = append(cards, "2 of Hearts")
	printSlice(cards)
}

func printSlice(cards []string) {

	for i, card := range cards {
		fmt.Println(i, card)
	}
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
