package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Welcome to channels session !!")


	myCh := make(chan int, 2)
 	wg:= &sync.WaitGroup{}
	

	// myCh <- 5
	// fmt.Println(<-myCh)

	// fmt.Println(<-myCh)
	// myCh <- 5

	wg.Add(2)

	//<-chan means Read only, R only channels can't close a channel

	go func(ch <-chan int, wg *sync.WaitGroup){ 
		val, isChannelOpen := <-myCh

		fmt.Println(val)
		fmt.Println(isChannelOpen)
		fmt.Println(<-myCh)
		wg.Done()
	}(myCh, wg)

	// chan<- writes or sends data into the channel
	
	go func(ch chan<- int, wg *sync.WaitGroup){
		myCh <- 0
		close(myCh)
		// myCh <- 5
		// myCh <- 6
		
		wg.Done()
	}(myCh, wg)
	
	wg.Wait()
	
}
