package main

import "fmt"

func main() {
	fibonacci := []int{5}
	returnData := []int{2}

	for i := 0; i < 10; i++ {
		var return2, totalNumber = Sum(returnData[i], fibonacci[i])
		returnData = append(returnData, return2)
		fibonacci = append(fibonacci, totalNumber)
	}
	fmt.Println(fibonacci)
}

func Sum(num1 int, num2 int) (int, int) {
	var sumNumber = num1 + num2
	return num2, sumNumber
}
