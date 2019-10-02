package main

import (
	"fmt"
	fibonacci "goclass/fibonacci"
)

func main() {
	f := fibonacci.Init{
		5, 7, 500
	}
	fmt.Println(f.FindNumberOfFibonacci)
}