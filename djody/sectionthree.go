package main

import (
	f "fmt"
	"fibonacci"
)

func Hitung(i int, j int){
	f.Println(i)
	f.Println(j)
	
	for index := 2; index < 10; index++{
		fij := i + j
		
		f.Println(fij)
		
		i = j
		j = fij
	}
}

func main(){
	i := 5
	j := 7
	
	f.Println("Barisan Fibonacci : ")
	fibonacci.Hitung(i, j)
}