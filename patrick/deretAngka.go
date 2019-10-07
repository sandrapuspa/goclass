package patrick

import "fmt"

func DeretAngka() {

	deret := []int{1}
	nilaiTambah := []int{3}

	for i := 0; i < 10; i++ {
		deret = append(deret, sum(deret[i], nilaiTambah[i]))
		nilaiTambah = append(nilaiTambah, plus2(nilaiTambah[i]))
	}

	fmt.Println(deret)

}

func sum(num1 int, num2 int) int {
	return num1 + num2
}

func plus2(number int) int {
	return number + 2
}
