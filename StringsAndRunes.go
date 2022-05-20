package main

import (
	"fmt"
	"unicode/utf8"
)


func main(){
	const s = "สวัสดี"

	fmt.Println("Len:", len(s))

	for i := 0; i <len(s); i++{
		fmt.Println("%x ", s[i])
	}
	fmt.Println()

	fmt.Println("Rune cont:", utf8.RuneCountInString(s))


	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}


	fmt.Println("\nUsing DecodeRuneInString")

	for i, w:= 0, 0; i< len(s); i+=w{
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width
		examineRune(runeValue)
	}
}

func examineRune(r rune) {
	if r=='t'{
		fmt.Println("found fee")
	}else if r=='ส'{
        fmt.Println("found so sua")
    }
	
}
