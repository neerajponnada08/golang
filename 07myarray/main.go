package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Println("FruitLst is: ", fruitList)
	fmt.Println("FruitLst length is: ", len(fruitList))


	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("Veggy list is: ", len(vegList))

	
}
