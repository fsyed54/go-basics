package main

import "fmt"

func main() {
	// var colors map[string] string

	colors := make(map[string] string)

	// colors := map [string] string{ // key value 
	// 	"red": "#ff0000",
	// 	"green": "#4bf745",
	// }

	colors["White"] = "#fffff"

	delete(colors, "White")

	fmt.Println(colors)
}