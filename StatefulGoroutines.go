package main

// each piece of data owned by exactly 1 goroutine

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// struct for read request
type readOp struct{
	key int
	resp chan int // this channel is used by owner to send back response
}

type writeOp struct{
	key int
	val int
	resp chan bool
}

func main(){
	var readOps uint64
	var writeOps uint64

	reads := make(chan readOp)
	writes := make(chan writeOp)

	go func(){
		// this routine owns the memory slots
		var state = make(map[int]int)
		for {
			select{
			case read:= <-reads:
				read.resp <-state[read.key] // send response via the given resp channel
			case write := <-writes:
				state[write.key] = write.val // write the val into states
				write.resp <-true	//
			}
		}
	}()

	for r:=0; r<100;r++{
		go func(){
			for{
				// create read request
				read := readOp{
					key: rand.Intn(5),
					resp: make(chan int), // I wonder if this would lead to heavy overhead
				}
				// send the request via reads channel
				reads <- read
				<-read.resp // wait for the response
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w:=0; w<10; w++{
		go func(){
			for{
				// create write request
				write := writeOp{
					key: rand.Intn(5),
					val: rand.Intn(100),
					resp: make(chan bool),
				}
				writes <- write // send via writes.channel
				<-write.resp // wait for response
				atomic.AddUint64(&writeOps,1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}
