package main
 
import "fmt"

import "goclass/andi/library"
 
func main(){    
    fmt.Print("TASK 1: Barisan Fibonacci : ")
    library.HitungFibonacci(10,5,7,0)
    fmt.Println()
    fmt.Print("TASK 2: Menampilkan angka prima start : 3 (10 angka) = ")
    library.Hitungprima(3,35)
    fmt.Println()
    fmt.Print("TASK 3: Hitung deret : 3 (10 angka) = ")
    library.Hitungderet(1,10,3)
}

