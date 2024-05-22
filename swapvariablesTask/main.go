package main

import "fmt"

func swapContent(listObj []int){ //Swap Content Condition
	for i, j := 0, len(listObj)-1; i < j; i, j = i+1, j-1{
		listObj[i], listObj[j] = listObj[j], listObj[i]
	}
}

func main() {
	listObj := []int {1,2,3,4,5,6,7} //Slice 
	swapContent(listObj)
	fmt.Println(listObj)
}

