package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("go channels");
	var myCha= make( chan int,1);

	// myCha<-5;
	// fmt.Println(<-myCha)

	wg:=& sync.WaitGroup{};

	wg.Add(2);
	go func (ch chan int , wg *sync.WaitGroup){
		// close(myCha)
		fmt.Println(<-myCha);
		fmt.Println(<-myCha);
		wg.Done();
	}(myCha,wg);
	go func (ch chan int , wg *sync.WaitGroup){
		// close(myCha)
		myCha<-5;
		myCha<-6;
		// close(myCha)
		myCha<-7;
		wg.Done();
	}(myCha,wg);

	wg.Wait();
}