package main

import "fmt"

func main() {
	fmt.Println("\nWelcome to functions!!\n")
	greeter()
	greeter2()

	result := adder(3,5)
	fmt.Println("Result is: ",result)

	
	proresult, Mymessage := proAdder(2,5,6,7,9)
	fmt.Println("Pro adder output: ",proresult)
	fmt.Println("Pro message is: ",Mymessage)

}

func adder(i1 int, i2 int) int{
	return i1+i2
}

func greeter() {

	fmt.Println("Namasthe from Golang!!\n")
}

func greeter2() {
	fmt.Println("Another method\n")
}

func proAdder(values ...int) (int, string){
	total := 0
	for _, val := range values{
		total += val
	}
	return total, "Hello return"
}

