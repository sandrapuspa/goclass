package main

import (
	"fmt"
)

func main() {
	fmt.Print("Test Exercise barisan angka Fibonacci : ")
	fib(5, 7, 10)

}
func fib(a, b, c int) {
	n1 := a
	n2 := b
	nall := []int{a, b}

	nextval := 0
	for a := 1; a <= c-2; a++ {
		nextval = n1 + n2
		n1 = n2
		n2 = nextval
		nall = append(nall, nextval)
	}

	for i := 0; i < len(nall); i++ {
		fmt.Print(nall[i], " ")
	}

}
