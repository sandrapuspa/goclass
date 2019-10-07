package main

import "fmt"

const maxFibo = 300

func fibo(a int, b int) int {
	c := a + b
	fmt.Print(c, " ")
	if c > maxFibo {
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
