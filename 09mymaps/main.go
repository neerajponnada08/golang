package main

import "fmt"

func main() {
	fmt.Println("Welcome to Maps of GoLang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("Languages: ",languages)
	fmt.Println("JS shorts for: ",languages["JS"])

	delete(languages, "RB")
	fmt.Println("Deleted RUby: ",languages)

	//loops are interesting in golang

	for key,value := range languages{
		fmt.Printf("For key %v, value is %v\n",key,value)

	}
}
