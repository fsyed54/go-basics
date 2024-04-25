package main

import "fmt"

func main() {
	colors := map [string] string{ // key value 
		"red": "#ff0000",
		"green": "#4bf745",
		"white": "#fffff",
	}

	printMap(colors)
}

func printMap(c map[string] string) { //c- color arg
	for color, hex := range c { //Color- key, Hex- value
		fmt.Println("Hex code for", color, "is", hex)
	}
}


// Map vs Struct

// Map: All keys and values must be of same type
// Struct: Values can be of different type

// Map: All the different keys are indexed- Iterate over different key value pairs
// Struct: Cannot iterate over

// Map: Reference type
// Struct: Value type
