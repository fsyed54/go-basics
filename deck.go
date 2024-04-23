package main

import "fmt"

//Create a new type of deck
//Which is a slice of strings

type deck []string

func newDeck() deck {
	cards:= deck{}

	cardSuits:= []string {"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues:= []string {"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits{
		for _, value := range cardValues{
			cards = append(cards, value+ " of " +suit)
		}
	}
	return cards
}

func (cards deck) print(){
	for i, card:= range cards{
		fmt.Println(i, card)
	}
}

// Slice Range Syntax

// abc[startIndexIncluding : upToNotIncluding]
// can be written as 
// abc[0:3] also abc[:3]
// similar abc[3:]