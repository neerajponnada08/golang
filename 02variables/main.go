package main

import "fmt"

const LoginToken string = "djfhrutydhsjk" //First letter capital implies public variable

func main() {
	var username string = "Neeraj"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n\n",username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n\n",isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n\n",smallVal)

	var smallFloat float32 = 255.3456785632467463838099
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n\n",smallFloat)

	var largeFloat float64 = 255.3456785632467463838099
	fmt.Println(largeFloat)
	fmt.Printf("Variable is of type: %T \n\n",largeFloat)

	//default values and some aliases

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n\n",anotherVariable)

	var anotherVariable1 string
	fmt.Println(anotherVariable1)
	fmt.Printf("Variable is of type: %T \n\n",anotherVariable1)

	//implicit type

	var website = "learncodeonline.com"
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T \n\n", website)

	//no var keyword

	noofitems := 100.0 // Doesn't work outside the functions
	fmt.Println(noofitems)
	fmt.Printf("Variable of type: %T \n\n", noofitems)


	//printing public variable from inside
	fmt.Println(LoginToken)
	fmt.Printf("Variable of type: %T \n\n", LoginToken)
}

