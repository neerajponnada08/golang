package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("LCO Web request!!")

	response, err := http.Get(url)

	if err!=nil {
		panic(err)
	}

	fmt.Printf("Response is of type: %T \n", response)

	defer response.Body.Close() //coller's responsibility to close the connection

	dataBytes, err := ioutil.ReadAll(response.Body)

	if err!=nil {
		panic(err)
	}

	content := string(dataBytes)
	fmt.Println("Content is ",content)
}
