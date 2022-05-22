package main

import (
	"fmt"
	"sort"
)

// byLength is alias for []string
type byLength []string

// sort interface: Len, Less, Swap

func (s byLength) Len() int{
	return len(s)
}

func (s byLength) Swap(i, j int){
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i,j int) bool{
	return len(s[i]) < len(s[j])
}

func main(){
	fruits := []string{"peach", "banana", "kiwi"}
	// byLength like a wrapper?
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
