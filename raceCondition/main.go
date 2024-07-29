package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("started with race conditon");

	var score []int;
	wg:=& sync.WaitGroup{};
	mut:= & sync.Mutex{};
	wg.Add(3);
	go func (wg *sync.WaitGroup,mu *sync.Mutex){
		fmt.Println("one R")
		mu.Lock()
		score = append(score, 1)
		mu.Unlock()
		defer wg.Done()
	}(wg,mut)
	go func (wg *sync.WaitGroup,mu *sync.Mutex){
		fmt.Println("two R")
		mu.Lock()
		score = append(score, 2)
		mu.Unlock()
		defer wg.Done()
	}(wg,mut)
	go func (wg *sync.WaitGroup,mu *sync.Mutex){
		fmt.Println("three R")
		mu.Lock()
		score = append(score, 3)
		mu.Unlock()
		defer wg.Done()
	}(wg,mut)

	wg.Wait()

	fmt.Println(score);
}