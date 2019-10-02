package main

import (
	"fmt"
	putra "github.com/ardhiansyahputra/goclass/putra"
)

func main() {
	f := fibonacci.Init{
		5, 7, 500
	}
	fmt.Println(f.FindNumberOfFibonacci)
}