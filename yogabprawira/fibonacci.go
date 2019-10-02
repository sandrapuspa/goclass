package main

import "fmt"

func fibo(a int, b int) int {
	c := a + b
	fmt.Print(c, " ")
	if c > 300 {
		return 0
	}
	return fibo(b, c)
}

func main() {
	a := 5
	b := 7
	fmt.Print(a, " ", b, " ")
	fibo(a, b)
}
