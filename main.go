package main

import (
	"fmt"
	"weather-checker/getall"
	"sync"
	"time"
)


var wg sync.WaitGroup
func main(){
	
	start := time.Now()
	wg.Add(5)
	go getall.Getall(&wg,"Moscow" )
	go getall.Getall(&wg,"Berlin")
	go getall.Getall(&wg,"Brasov")
	go getall.Getall(&wg,"Tokyo")
	go getall.Getall(&wg,"Odessa")
	wg.Wait()
	timenow := time.Since(start)
	fmt.Println(timenow)
}