package main

import "fmt"


// a function that returns an anonymous function
func intSeq() func() int{
	i := 0

	// each instance keeps its own closure of variables
	return func() int{
		i++
		return i
	}
}

func main(){
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())

}
