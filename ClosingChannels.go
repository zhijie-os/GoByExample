package main

import "fmt"

func main(){
	jobs := make(chan int, 5)
	done := make(chan bool)


	go func(){
		for{
		// j would store the actual job, and more would either be false if jobs has been closed and all values in the channel have already been received
			j, more := <-jobs
			if more{
				fmt.Println("received job", j)
			} else{
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++{
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	// wait for worker
	<-done
}
