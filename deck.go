package main

import "fmt"

// Create a type of 'decl'
// which is a slice of strings

type deck []string

func (d deck) print() {

	for i, card := range d {
		fmt.Println(i, card)
	}
}
