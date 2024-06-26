package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

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

func (d deck) print(){
	for i, card:= range d{
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]

}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err !=nil {
		// If something goes wrong
		// Option #1- log the error and return a call to newDeck()
		// Option #2- log the error and entirely quit the program
		fmt.Println("Error:",err)
		os.Exit(1)
	}
	s:= strings.Split(string(bs), ",") //Splitting a slice of strings package
	return deck(s)
}

func (d deck) shuffle(){
	for i := range d {
		newPosition := rand.Intn(len(d) - 1) //length of slice --> len(d)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}