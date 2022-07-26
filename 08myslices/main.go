package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices in GoLang")

	var fruitList = []string{"Apple","Tomato","Peach"}

	fmt.Printf("Type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Guava", "Pomegranate")

	fmt.Println("FruitList: ",fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScore := make([]int, 4)

	highScore[0] = 234
	highScore[1] = 945
	highScore[2] = 465
	highScore[3] = 867
	// highScore[4] = 777  //Result error

	highScore = append(highScore, 555, 666, 777)

	fmt.Println("Highscore: ",highScore)

	//Sorting
	sort.Ints(highScore)
	fmt.Println("Highscore Sorted: ",highScore)
	fmt.Println(sort.IntsAreSorted(highScore))

	var courses = []string{"reactjs","javascript","swift","python","ruby"}
	fmt.Println("courses: ",courses)

	//remove the values using append
	var index int=2
	courses = append(courses[:index],courses[index+1:]...)
	fmt.Println("Courses: ",courses)
}
