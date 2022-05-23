package main

import (
	"fmt"
	"time"
)

func main(){
	now := time.Now()
	fmt.Println(now)

	fmt.Println(now.Unix())
	//fmt.Println(now.UnixMilli())
	fmt.Println(now.UnixNano())

	// convert integer seconds or nanoseconds since the epoch into the corresponding time
	fmt.Println(time.Unix(now.Unix(),0))
	fmt.Println(time.Unix(0,now.UnixNano()))
}
