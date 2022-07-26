package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("RACE Condition !!")

	wg := &sync.WaitGroup{}

	mut := &sync.RWMutex{}

	
	var score = []int{0}
	

	wg.Add(4)
	// wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.RWMutex){
		fmt.Println("One R")
		m.Lock()
		score  = append(score, 1)
		m.Unlock()
		wg.Done()
	}(wg, mut)

	// wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.RWMutex){
		fmt.Println("Two R")
		m.Lock()
		score  = append(score, 2)
		m.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex){
		fmt.Println("Three R")
		m.Lock()
		score  = append(score, 3)
		m.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex){
		fmt.Println("Four R")
		m.RLock()
		fmt.Println(score )
		m.RUnlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	
	fmt.Println(score)

}
