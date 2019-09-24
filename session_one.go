package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("response 500")
	fmt.Printf("%T = %v\n", err, err)

	em := "hello"
	fmt.Printf("%T = %v\n", em, em)

	fl := 14.0 / 3.0
	fmt.Println(fl)
	fmt.Printf("%.2f", fl)

	const nama string = "David Maulana"
	fmt.Println(nama)

	// mapType := make(map[int]string)
	// mapType[0] = "string1"
	// mapType[100] = "string100"

	mapType := map[int]string{
		0:   "string1",
		100: "string100",
	}

	for key, value := range mapType {
		fmt.Println(key, "===", value)
	}

	mapString := make(map[string][]string)
	mapString["patrick"] = []string{"Strudents", "intern"}

	for key, value := range mapString {
		fmt.Println(key, "===", value)
	}

	sliceString := []string{"employee", "sde"}
	for _, value := range sliceString {
		fmt.Println("===", value)

	}

	doubleArray := [2][4]int{
		[4]int{1, 2, 3, 4},
		[4]int{5, 6, 7, 8},
	}
	for key, value := range doubleArray {
		fmt.Println(key, "===", value)
	}

	
	data := [10]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	var aSlice, bSlice []string

// operasi standar
aSlice = data[:3] // sama dengan aSlice = data[0:3], hasilnya a,b,c
fmt.Println("aSlice", aSlice)

aSlice = data[5:] // sama dengan aSlice = data[5:10], 10 merupakan panjang dari data, hasilnya f,g,h,i,j

fmt.Println("aSlice", aSlice)
aSlice = data[:] // sama dengan aSlice = data[0:10], hasilnya menampikan semua elemen

// re-slice

fmt.Println("aSlice", aSlice)
aSlice = data[3:7] // mulai dari index=3 sampai kapasitas=7, hasilnya d,e,f,g

fmt.Println("aSlice", aSlice)
bSlice = aSlice[1:3] // mulai dari index=1 sampai kapasitas=3, hasilnya e,f
bSlice = aSlice[:3] // mulai dari index=0 sampai kapasitas=3, hasilnya d,e,f
bSlice = aSlice[0:5] // slice dapat bertambah melebihi kapasitasnya karena slice hanya mng-copy data (nanti dipelajari lebih lanjut), hasilnya d,e,f,g,h
bSlice = aSlice[:] // bSlice memiliki elemen seperti aSlice, hasilnya d,e,f,g
fmt.Println("bSlice", bSlice)


slice := make([]int, 2, 10)
slice = append(slice, 8)
slice[0]=10
fmt.Println(slice) // [8]
for key, value := range slice {
	fmt.Println(key, "===", value)
}

slices := make([]int, 0, 10)
slices = slices[0:10]
slices[7] = 8
slices = append(slices, 100)
fmt.Println(slices)

scores := []int{1, 2, 3, 4, 5}
slicess := scores[2:4]
slicess[0] = 999

fmt.Println("slice: ", slicess)
fmt.Println("scores: ", scores)



map1 := make(map[string]string)
map2 := make(map[string]string)
map1["one"] = "Satu"

fmt.Println("map1: ", map1)
map2 = map1
fmt.Println("map2: ", map2)
map1["one"] = "Sepuluh" 
fmt.Println("map1after: ", map1)// maka map2["one"] adalah Sepuluh
fmt.Println("map2after: ", map2)
}

func testFunc() {
	number := 0

testLabel:
	fmt.Println(number)
	number++
	goto testLabel
}
