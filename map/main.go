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