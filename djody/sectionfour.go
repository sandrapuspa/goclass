package main

import (
	f "fmt"
)

func main(){
	input := 3
	total := 0
	
	f.Print("Bilangan Prima : ");

	for i := input; i < 100; i++{
		k := 0
		
		for j := 1; j < i; j++{
			if i % j == 0{
				k++
			}
		}
		
		if k < 2{
			if total < 10 {
				f.Print(i, " ")
				total++
			}
		}
	}
}