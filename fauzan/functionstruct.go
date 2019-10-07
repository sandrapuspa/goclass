package main

import (
	"fmt"
	"os"
)

var opts = []struct {
	nama, alamat, pekerjaan, alasan string
}{
	{"Fauzan Rabbani", "Jaksel", "Barista", "Mengisi Waktu Luang"},
	{"Rongonald", "Jakut", "Atlet", "Mengisi Waktu Luang"},
}
var arg int

func main() {
	arg = os.Args[1]
	b := opts[arg]
	fmt.Println("hello")
	fmt.Println(b)
	fmt.Println(arg)

}
