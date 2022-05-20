package main

import "fmt"

func main(){
	// channels are the pipes that connect concurrent goroutines

	// create new channel via make(chan val-type)
	messages := make(chan string)
	// send a value via channel <- syntax


	// By default sends and receives block until both the sender and receiver are ready. 
	go func() { messages <- "ping"}()

	msg := <-messages
	fmt.Println(msg)
}
