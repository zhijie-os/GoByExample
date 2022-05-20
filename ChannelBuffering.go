package main

import "fmt"

func main(){
	// make a channel of strings buffering up to 2 values
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
