package main

import (
	"fmt"
	"strings"
)

func main(){
	//Create and Initialize strings
	string1 := "Welcome to Miracle"
	string2 := "Go Interview Preparation"

	//Check contains
	res1 := strings.Contains(string1, "Miracle")
	res2 := strings.Contains(string2, "Hello")

	//Printing results
	fmt.Println("Is 'Miracle' present in string 1:", res1)
	fmt. Println("Is 'Go' present in string 2:", res2)

}