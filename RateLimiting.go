package main

import (
	"fmt"
	"time"
)


func main(){
	// create a channel to feed request
	requests := make(chan int, 5)
	for i:=1; i <= 5; i++{
		requests <- i	// feed requests 1 to 5
	}
	close(requests)	// no more requests

	// time limiter, 200ms
	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests{
		<-limiter // for each request, wait for tick
		fmt.Println("request", req, time.Now())
	}









	// we allow short bursts of requests in rate limiter
	// <-> buffering limiter channel
	burstyLimiter := make(chan time.Time, 3)


	// fill up the burstyLimiter
	for i:=0; i<3;i++{
		burstyLimiter <- time.Now()
	}


	// for every 200ms, try to add a new value
	go func(){
		for t := range time.Tick(200 * time.Millisecond){
			burstyLimiter <- t
		}
	}()

	// burstyRequests := make(chan int, 5)
	burstyRequests := make(chan int, 5)
	for i:= 1; i<=5; i++{
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests{
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}



}
