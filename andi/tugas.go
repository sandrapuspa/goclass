package main
 
import "fmt"

import "library"
 
func main(){    
    fmt.Print("TASK 1: Barisan Fibonacci : ")
    library.HitungFibonacci(10,5,7,0)
    fmt.Println()
    fmt.Print("TASK 2: Menampilkan angka prima start : 3 (10 angka) = ")
    hitungprima(3,35)
    fmt.Println()
    fmt.Print("TASK 3: Hitung deret : 3 (10 angka) = ")
    hitungderet(1,10,3)
}

/*func fibonacci(n int, t1 int, t2 int, nextnya int) {
    for i:=1;i<=n;i++ {
        if(i==1){
            fmt.Print(" ",t1)
            continue
        }
        if(i==2){
            fmt.Print(" ",t2)
            continue
        }
        nextnya = t1 + t2
        t1 = t2
        t2 = nextnya
        fmt.Print(" ",nextnya)
    }
}*/

func hitungprima(awal int, akhir int) {
    for i:=awal;i<=akhir;i++ {
        a := 0
        for j:=1;j<=i;j++ {
            if(i % j == 0){
                a++
            }
        }
        if(a==2){
            fmt.Print(" ",i)
        }
    }
}

func hitungderet(start int, n int, mp int) {
    var i int 
    count := 0
    next := start
    for i = start; i <= 190; i++ {

        if count == n {
            break
        }

        fmt.Print(next, " ")
        next = next + mp
        mp = mp + 2
        count = count + 1
    }
}