package main

import (
	"fmt"
	fibonacci "goclass/fibonacci"
)

// func fibonacci(val1, val2, limit int) []int {
// 	var result = []int{val1, val2}

// 	for {
// 		val3 := val1 + val2
// 		val1 = val2
// 		val2 = val3

// 		if (val3 > limit) {
// 			break
// 		}

// 		result = append(result, val3)
// 	}

// 	return result
// }

func main() {
	f := fibonacci.Init{
		5, 7, 500
	}
	fmt.Println(f.FindNumberOfFibonacci)
}