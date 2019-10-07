package main

import (
	"fmt"

	putra "goclass/putra"
)

func main() {
	f := putra.Init{
		A:     5,
		B:     7,
		Limit: 50,
	}
	fmt.Println(f.FindNumberOfFibonacci)
}
