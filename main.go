package main

import (
	"fmt"

	putra "github.com/ardhiansyahputra/goclass/putra"
)

func main() {
	f := putra.Init{
		A:     5,
		B:     7,
		Limit: 50,
	}
	fmt.Println(f.FindNumberOfFibonacci)
}
