package main

import (
	"flag"
	"fmt"
)

func main() {

	// syntax (name, default, description)
	wordPtr := flag.String("word", "foo", "a string")

	numbPtr := flag.Int("numb", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	// different syntax
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// need to parse the flags
	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)

	// flag.Args() would give back the tail 
	fmt.Println("tail:", flag.Args())
}
