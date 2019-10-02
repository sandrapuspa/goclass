package main

import "fmt"

func main() {
	start := 1
	for i := start; i < start+10; i++ {
		fmt.Print(i*i, " ")
	}
}
