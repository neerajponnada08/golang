package main

import "fmt"

func main() {

	defer fmt.Println("Hello world")
	defer fmt.Println("one")
	myDefer()

	defer fmt.Println("two")
	fmt.Println("Welcome to differ !!")
	
	
	

	
}
func myDefer(){
	for i:=0;i<5;i++ {
		defer fmt.Print(i)
	}
	fmt.Println()
}