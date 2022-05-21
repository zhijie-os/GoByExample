package main

import "fmt"
// pings = to send
func ping(pings chan<- string, msg string){
	pings <- msg
}

//		pings = to receive, pongs to send
func pong(pings <-chan string, pongs chan<- string){
	msg:= <-pings
	pongs <- msg
}

func main(){
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	// send msg through pings
	ping(pings, "passed message")
	// receive from pings and send to pongs
	pong(pings, pongs)
	// receive from pongs and printout
	fmt.Println(<-pongs)
}
