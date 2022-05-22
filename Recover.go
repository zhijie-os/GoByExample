package main

import "fmt"

func mayPanic(){
	panic("a problem")
}

func main(){
	// defer would be executed after the execution of the program
	defer func(){
		if r:=recover(); r!=nil{
			fmt.Println("Recoverd, Error:\n", r)
		}
	}()

	mayPanic() // program terminates here
	// never reached
	fmt.Println("After mayPanic()")
}
