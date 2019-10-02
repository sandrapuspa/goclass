package main

import "fmt"

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	start := 3
	i := start
	n := 0
	for {
		if n >= 10 {
			break
		}
		if isPrime(i) {
			fmt.Print(i, " ")
			n++
		}
		i++
	}
}
