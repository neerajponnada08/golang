package main

import "fmt"

func main() {
	fmt.Println("Welcome to class on pointers")

	// var ptr *int
	// fmt.Println("Value of ptr, ",ptr)

	myNumber := 23

	//If reference is the case we can use &
	var ptr = &myNumber 

	fmt.Println("Value of ptr , ",ptr)
	fmt.Println("Value of actual ptr , ",*ptr)

	*ptr = *ptr+2
	fmt.Println("New value, ",myNumber)


}
