package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)
var wg sync.WaitGroup;
var mut sync.Mutex
var data []string;
func main() {

	websites := []string{
		"https://go.dev",
		"https://fb.com",
		"https://vipinnotes.onrender.com",
		"https://github.com",
		"https://google.com",
	}

	for _,web:= range websites{
		go getDetailsOfWebsite(web);
		wg.Add(1)
	}

	wg.Wait();
}

func getDetailsOfWebsite(website string) {
	res,err:=http.Get(website);

	if err!=nil {
		log.Fatal(err);
	}else{
		mut.Lock()
		data=append(data, website);
		mut.Unlock()
		
		fmt.Printf("%v is the statuscode for %s \n",res.StatusCode,website);
	}
	defer wg.Done()
}