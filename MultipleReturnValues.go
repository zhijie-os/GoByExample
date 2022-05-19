package main

// vim
// i: insert, o: insert at new line
// shift+o: insert at new line above
// shift+i: insert at the beginning of the line 
// shift+a: insert at the end the line 



// down arrow: move cursor one line down
// up arrow: move cursor one line above
// n, down arrow: move cursor n lines down
// n, up arrow: move cursor one line above


// k<-> up arrow
// j<-> down arrow 
// h<-> left arrow 
// l<-> right arrow



// u: undo
// ctrl+r: redo
// n+u: undo n 
// n+ctrl+r: redo n

// v: enter visual mode
// d: delete selected content
// y: yanking, copy
// p: paste
// dd: delete whole line
// yy: copy whole line
// D: delete the rest of the line
// C: change the rest of the line
// r: replace
// w: jump to next word
// b: jump to previous word
// dw: delete a word,  2dw: delete 2 words
// diw: delete inner word
// ciw: change in a word
// e: jump to the end of the word
// 0: jump to the beginning of the line
// $: jump to the end of the line
import "fmt"

// %: jump to the closing/opening bracket
// d%: delete everything include and within the bracket

// ==: auto indent

func vals() (int, int){
	return 3,7
}

func main(){

	a,b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c:=vals()
	fmt.Println(c)

}




// zz: center the cursor
// /pattern to search
// n: move to the next occurence
// N: move to the previous occurrence

// gg: go to the beginning of the file
// G: go to the end of the file
// gg=G: indent entire file

// >> shift right
// << shift left

// m + a = mark point a
// ' + a = jump to a

// .: last command


