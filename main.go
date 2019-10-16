package main

import (
	"fmt"
	// "runtime"
)

func sum(arrayInt []int, channelInt chan int) {
	total := 0
	for _, value := range arrayInt {
		// fmt.Println(value)
		
			total += value

			// runtime.Gosched()
	}
	fmt.Println("total:", total)

	channelInt <- total // send total to channelInt
}

func fibonacci(length int, channel chan int) {
	fmt.Println("Length:", length)
	number1, number2 := 1, 1
	for i := 0; i < length; i++ {
			channel <- number1
			number1, number2 = number2, number1+number2
	}
	close(channel)
}

func main() {
	// SessionFive()

	// arrayInt := []int{7, 2, 8, -9, 4, 0}

	// channelInt := make(chan int,1)
	// go sum(arrayInt[:len(arrayInt)/2], channelInt)
	// go sum(arrayInt[len(arrayInt)/2:], channelInt)
	// go sum(arrayInt[len(arrayInt)/3:], channelInt)
	// channelInt <- 1
	// result1 := <- channelInt
	// channelInt <- 2
	// result2 := <-channelInt // receive from channelInt

	// fmt.Println(result1, result2, result1+result2)

	kelas := [10]int{}
	channel := make(chan int, 10)
	go fibonacci(cap(channel), channel)
	fmt.Println("CLOSED", cap(kelas))
	for i := range channel {
			fmt.Println(i)
	}

	// channel := make(chan int, 2)
	// channel <- 1
	// channel <- 2
	// fmt.Println(<-channel)
	// fmt.Println(<-channel)
}