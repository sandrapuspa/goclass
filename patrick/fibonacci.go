package patrick

import "fmt"

func Fibonacci() {
	fibonacciSlice := []int{5}
	returnData := []int{2}

	for i := 0; i < 10; i++ {
		var return2, totalNumber = Sum(returnData[i], fibonacciSlice[i])
		returnData = append(returnData, return2)
		fibonacciSlice = append(fibonacciSlice, totalNumber)
	}
	fmt.Println(fibonacciSlice)
}

func Sum(num1 int, num2 int) (int, int) {
	var sumNumber = num1 + num2
	return num2, sumNumber
}
