package main

import "fmt"

func PrintAnything[T any](thing T) {
  fmt.Println(thing)
}

func main(){
	PrintAnything(2)
}
