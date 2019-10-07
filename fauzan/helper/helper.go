package helper

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var opts = []struct {
	nama, alamat, pekerjaan, alasan string
}{
	1: {"Fauzan Rabbani", "Jaksel", "Barista", "Mengisi Waktu Luang"},
	{"Steven", "Jakut", "Atlet", "Mengisi Waktu Luang"},
	{"Williams", "Jakbar", "Atlet", "Mengisi Waktu Luang"},
	{"Rudy", "Jaktim", "Atlet", "Mengisi Waktu Luang"},
	{"Claudio", "Jakpus", "Atlet", "Mengisi Waktu Luang"},
	{"Robert", "Lampung", "Atlet", "Mengisi Waktu Luang"},
	{"Alfredo", "Bengkulu", "Atlet", "Mengisi Waktu Luang"},
}

func Absen() {
	//arg := os.Args[1]

	//argg, _ := strconv.Atoi(arg)
	//fmt.Println(len(opts))
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan No Absen: ")
	// Scans a line from Stdin(Console)
	scanner.Scan()
	// Holds the string that scanned
	text := scanner.Text()

	argg, _ := strconv.Atoi(text)
	if argg >= len(opts) {
		fmt.Println("Tidak Ada No Absen ", argg)
	} else {
		b := opts[argg]

		fmt.Println("Nama : ", b.nama)

		fmt.Println("Alamat : ", b.alamat)
		fmt.Println("Pekerjaan : ", b.pekerjaan)
		fmt.Println("Alasan Mengikuti Kelas Golang : ", b.alamat)
	}
	return

}
func Fib(a, b, c int) {
	n1 := a
	n2 := b
	nall := []int{a, b}

	nextval := 0
	for a := 1; a <= c-2; a++ {
		nextval = n1 + n2
		n1 = n2
		n2 = nextval
		nall = append(nall, nextval)
	}

	for i := 0; i < len(nall); i++ {
		fmt.Print(nall[i], " ")
	}

}
func Inputdata() {
	// To create dynamic array
	//arr := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter Name: ")
	// Scans a line from Stdin(Console)
	scanner.Scan()
	// Holds the string that scanned
	ninput := scanner.Text()
	fmt.Print("Enter Alamat: ")

	scanner.Scan()
	ainput := scanner.Text()

	fmt.Print("Enter Pekerjaan: ")

	scanner.Scan()
	pinput := scanner.Text()
	fmt.Print("Enter Alasan: ")

	scanner.Scan()
	alinput := scanner.Text()
	/*for {
		fmt.Print("Enter Text: ")
		// Scans a line from Stdin(Console)
		scanner.Scan()
		// Holds the string that scanned
		text := scanner.Text()
		if len(text) != 0 {
			fmt.Println(text)
			arr = append(arr, text)
		} else {
			break
		}

	}*/
	// Use collected inputs
	//opts = append(opts, struct{ ninput, ainput, pinput, alinput string })
	fmt.Println("Nama ", ninput)
	fmt.Println("Alamat", ainput)
	fmt.Println("Pekerjaan", pinput)
	fmt.Println("Alasan", alinput)
}
