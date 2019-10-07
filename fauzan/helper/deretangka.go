package helper

import "fmt"

func Prima(start, n int) {
	var g, a, i int

	fmt.Print("Barisan Angka Prima -> ")
	count := 0
	for i = start; i <= 200; i++ {
		if count == n {
			fmt.Print("\n")
			break
		}
		a = 0
		for g = 1; g <= i; g++ {
			if i%g == 0 {
				a++
			}
		}

		if a == 2 {
			fmt.Print(i, " ")
			count = count + 1
		}

	}
}

func Deret(start, n, mp int) {
	var i int

	fmt.Print("Deret Angka -> ")
	count := 0
	next := start
	for i = start; i <= 200; i++ {

		if count == n {
			break
		}

		fmt.Print(next, " ")
		next = next + mp
		mp = mp + 2
		count = count + 1
	}
}
