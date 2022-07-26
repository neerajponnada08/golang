package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcvome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for pizza: ")

	//comma ok syntax or error ok syntax

	input, _ := reader.ReadString('\n')
	fmt.Println("\nThanks for rating, ", input)
	fmt.Printf("Type of input is %T",input)

}
