package main

import (
	f "fmt"
)

func main(){
	input := 1
	start_add := 3
	add_multiplier := 2
	
	for i := 0; i < 10; i++{
		f.Print(input, " ")
	
		input += start_add
		start_add += add_multiplier
	}
}