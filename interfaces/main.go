package main

import "fmt"

type bot interface{
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	return "Hi there!"
}

func (sb spanishBot) getGreeting() string {
	// Same logic
	return "Hola!"
}

// Concrete types: map, struct, int, string, englishBot
// Interface type: bot

// Interfaces are not generic types
// Interfaces are implicit
// Interfaces are tough
// Interfaces are a contract to help us manage types