package main

import (
	"fmt"
	"math/rand/v2"
)

// Hello returns a greeting for the named person.
func main() {
	// Return a greeting that embeds the name in a message.

	FibonacciSeries(rand.IntN(20))
	//fmt.Println("jehu ")
	//fmt.Printf("Hello, Gopher %9f ", 0.567)
	//Wrong type or unknown verb: %!verb(type=value)
	//fmt.Printf("%d", "hi") //:        %!d(string=hi)
	//Too many arguments: %!(EXTRA type=value)
	//fmt.Printf("hi", "guys") //:      hi%!(EXTRA string=guys)
	//Too few arguments: %!verb(MISSING)
	//fmt.Printf("hi%d") //:            hi%!d(MISSING)
	//Non-int for width or precision: %!(BADWIDTH) or %!(BADPREC)
	//fmt.Printf("%*s", 4.5, "hi")  //:  %!(BADWIDTH)hi
	//fmt.Printf("%.*s", 4.5, "hi") //: %!(BADPREC)hi
	//Invalid or invalid use of argument index: %!(BADINDEX)
	//fmt.Printf("%*[2]d", 7) //:       %!d(BADINDEX)
	//fmt.Printf("%.[2]d", 7) //:       %!d(BADINDEX)

}

func FibonacciSeries(n int) {
	t1 := 0
	t2 := 1
	i := 0
	nextItem := t1 + t2
	fmt.Printf(" %d %d ", t1, t2)

	for i = 3; i <= n; i++ {
		fmt.Printf(" %d", nextItem)
		t1 = t2
		t2 = nextItem
		nextItem = t1 + t2
	}
}
