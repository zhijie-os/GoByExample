package main

import (
	"fmt"
	"sync"
	"time"
)


func worker(id int){
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main(){

	// WaitGroup is used to wait for all the goroutines launched here to finish
	var wg sync.WaitGroup


	for i:=1; i <= 5; i++{
		//
		wg.Add(1)

		i := i

		go func(){
			// A defer statement defers the execution of a function until the surrounding function returns.
			defer wg.Done() // this would be called when function return
			worker(i)
		}()
	}


	wg.Wait()
}
