package main

import (
	"fmt"
	"net/http"
	"sync"
	// "time"
)

var signals = [] string {"test"}
var wg sync.WaitGroup //pointers
var mut sync.Mutex //pointer

func main() {
	fmt.Println("Welcome to go routines !!")

	// go greeter("Hello")
	// greeter("World")

	websitelist := []string {
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3*time.Millisecond)
// 		fmt.Println(s)
// 	}
// }


func getStatusCode(endpoint string) {

	defer wg.Done()

	res, err := http.Get(endpoint)

	if err!=nil {
		fmt.Println("OOPS in endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for website is %s\n", res.StatusCode, endpoint)

	}

}