package main

import (
	"bufio"
	"fauzan/helper"
	"fmt"
	"os"
)

func main() {
	Menu()
}

func Menu() {
	//helper.Absen()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("1. Input Data")
	fmt.Println("2. Lihat Data")
	fmt.Print("Masukkan Pilihan: ")
	// Scans a line from Stdin(Console)
	scanner.Scan()
	// Holds the string that scanned
	text := scanner.Text()
	if text == "1" {
		helper.Inputdata()
	} else {
		helper.Absen()
	}
}
