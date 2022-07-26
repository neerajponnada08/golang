package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("Welcome to sending form data !!")

	PerformPostFormRequest()

}

func PerformGetRequest() {

}

func PerformPostJsonRequest() {

}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	//Form data

	data := url.Values{}
	data.Add("firstname","neeraj")
	data.Add("lastname","vardhan")
	data.Add("email","neeraj@gmail")

	response, err:=http.PostForm(myurl, data)

	if err!=nil {
		panic(err)
	}

	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}
