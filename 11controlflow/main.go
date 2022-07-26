package main

import "fmt"

func main() {
	fmt.Println("Welcome to control statements")

	loginCount  := 15
	var result string

	if loginCount<10{
		result = "Regular user"
	} else if loginCount>=10 && loginCount<20{
		result = "Moderate User"	
	} else{
		result = "Favorite user"
	}

	fmt.Println("Result is: ",result)

	
	if num := 3; num <10 {
		fmt.Println("Num is less that 10")
	}	else {
		fmt.Println("Num is NOT less than 10")
	}

	  
}
