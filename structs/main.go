package main

import "fmt"

type contactInfo struct{
	email string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName: "Party",
		contactInfo: contactInfo{
			email: "jim@gmail.com",
			zipCode: 94000,
		},
	}

	jim.updateName("Jimmy") //Type of person now
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) { //pointer type
	(*pointerToPerson).firstName = newFirstName //operator, manipulating the value the pointer is referencing to
}

func (p person) print() {
	fmt.Printf("%+v", p)
}